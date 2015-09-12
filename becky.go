package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"go/build"
	"hash/fnv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

// As asset.go generates the source code assets, and bundle.go
// generates asset.go, there's actually a cycle here, where first run
// might not update everything. That's rare enough that we don't care.
//
//go:generate -command asset go run asset.go
//go:generate asset -lib=false -var=assetDev -wrap=asset -- asset_dev.go
//go:generate asset -lib=false -var=assetNoDev -wrap=asset -- asset_nodev.go
//go:generate go run ./internal/bundle.go

var (
	flagVar  = flag.String("var", "", "variable name to use, \"_\" to ignore (default: file basename without extension)")
	flagWrap = flag.String("wrap", "", "wrapper function or type (default: filename extension)")
	flagLib  = flag.Bool("lib", true, "generate asset_*.gen.go files defining the asset type")
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s [OPTS] FILE..\n", prog)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Creates files FILE.gen.go and asset_*.gen.go\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if flag.NArg() > 1 && *flagVar != "" && *flagVar != "_" {
		log.Fatal("cannot combine -var with multiple files")
	}

	// Look up package names and import paths only once. Every dir
	// here is also guaranteed to have the auxiliary files already, to
	// avoid repeating work.
	packages := map[string]*build.Package{}

	for _, filename := range flag.Args() {
		dir, base := filepath.Split(filename)
		if dir == "" {
			dir = "."
		}

		pkg, err := getPkg(packages, dir)
		if err != nil {
			log.Fatal(err)
		}

		variable := *flagVar
		if variable == "" {
			variable = strings.SplitN(base, ".", 2)[0]
		}

		wrap := *flagWrap
		if wrap == "" {
			wrap = filepath.Ext(base)
			if wrap == "" {
				log.Fatalf("files without extension need -wrap: %s", filename)
			}
			// remove the dot
			wrap = wrap[1:]
		}

		if err := process(filename, pkg.Name, variable, wrap); err != nil {
			log.Fatal(err)
		}

	}
}

// autogen writes a warning that the file has been generated automatically.
func autogen(w io.Writer) error {
	// broken into parts here so grep won't find it
	const warning = "// AUTOMATICALLY " /* but not this file */ + "GENERATED FILE. DO NOT EDIT.\n\n"
	_, err := io.WriteString(w, warning)
	return err
}

func process(filename, pkg, variable, wrap string) error {
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	tmp, err := ioutil.TempFile(filepath.Dir(filename), ".tmp.asset-")
	if err != nil {
		return err
	}
	defer func() {
		if tmp != nil {
			_ = os.Remove(tmp.Name())
		}
	}()
	defer tmp.Close()

	in := bufio.NewReader(src)
	out := bufio.NewWriter(tmp)

	if err := autogen(out); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "package %s\n\n", pkg); err != nil {
		return err
	}
	if err := embed(variable, wrap, filepath.Base(filename), in, out); err != nil {
		return err
	}
	if err := out.Flush(); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	gen := filename + ".gen.go"
	if err := os.Rename(tmp.Name(), gen); err != nil {
		return err
	}
	tmp = nil
	return nil
}

func embed(variable, wrap, filename string, in io.Reader, out io.Writer) error {
	h := fnv.New64a()
	r := io.TeeReader(in, h)
	_, err := fmt.Fprintf(out, "var %s = %s(asset{Name: %q, Content: \"\" +\n",
		variable, wrap, filename)
	if err != nil {
		return err
	}
	buf := make([]byte, 1*1024*1024)
	eof := false
	for !eof {
		n, err := r.Read(buf)
		switch err {
		case io.EOF:
			eof = true
		case nil:
			// nothing
		default:
			return err
		}
		if n == 0 {
			continue
		}
		s := string(buf[:n])
		s = strconv.QuoteToASCII(s)
		s = "\t" + s + " +\n"
		if _, err := io.WriteString(out, s); err != nil {
			return err
		}
	}
	etag := `"` + base64.StdEncoding.EncodeToString(h.Sum(nil)) + `"`
	if _, err := fmt.Fprintf(out, "\t\"\", etag: %#q})\n", etag); err != nil {
		return err
	}
	return nil
}

func getPkg(packages map[string]*build.Package, dir string) (*build.Package, error) {
	if pkg, found := packages[dir]; found {
		return pkg, nil
	}

	pkg, err := loadPkg(dir)
	if err != nil {
		return nil, err
	}
	if *flagLib {
		if err := auxiliary(pkg.Dir, pkg.ImportPath, pkg.Name); err != nil {
			return nil, err
		}
	}
	packages[dir] = pkg
	return pkg, nil
}

func loadPkg(dir string) (*build.Package, error) {
	// if we don't resolve to absolute paths, the resulting import
	// path ends up as just "."
	if !filepath.IsAbs(dir) {
		if abs, err := filepath.Abs(dir); err == nil {
			dir = abs
		}
	}

	pkg, err := build.ImportDir(dir, 0)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

func auxiliary(dir, imp, pkg string) error {
	for filename, tmpl := range map[string]string{
		"asset_dev":   assetDev.Content,
		"asset_nodev": assetNoDev.Content,
	} {
		tmpl = strings.Replace(tmpl, "\npackage main\n", "\npackage "+pkg+"\n", 1)

		t, err := template.New("").Parse(tmpl)
		if err != nil {
			return err
		}

		tmp, err := ioutil.TempFile(dir, ".tmp.asset-")
		if err != nil {
			return err
		}
		defer func() {
			if tmp != nil {
				_ = os.Remove(tmp.Name())
			}
		}()
		defer tmp.Close()

		type data struct {
			ImportPath string
		}
		d := data{
			ImportPath: imp,
		}
		if err := autogen(tmp); err != nil {
			return err
		}
		if err := t.Execute(tmp, d); err != nil {
			return err
		}
		if err := tmp.Close(); err != nil {
			return err
		}
		gen := filepath.Join(dir, filename+".gen.go")
		if err := os.Rename(tmp.Name(), gen); err != nil {
			return err
		}
		tmp = nil
	}
	return nil
}
