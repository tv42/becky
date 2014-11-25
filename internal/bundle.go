// +build ignore

// bundle.go writes a copy of main.go with snippets inlined and build
// tag "ignore" set, to asset.go.
package main

import (
	"errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("bundle.go: ")

	if err := bundle(); err != nil {
		log.Fatal(err)
	}
}

var toBundle = map[string]struct{}{
	"becky.go":              struct{}{},
	"asset_nodev.go":        struct{}{},
	"asset_dev.go.gen.go":   struct{}{},
	"asset_nodev.go.gen.go": struct{}{},
}

func filter(fi os.FileInfo) bool {
	_, found := toBundle[fi.Name()]
	return found
}

func bundle() error {
	tmp, err := ioutil.TempFile(".", ".tmp.asset-")
	if err != nil {
		return err
	}
	defer func() {
		if tmp != nil {
			_ = os.Remove(tmp.Name())
		}
	}()
	defer tmp.Close()

	if _, err := io.WriteString(tmp, "// +build ignore\n\n"); err != nil {
		return err
	}

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, ".", filter, parser.ParseComments)
	if err != nil {
		return err
	}
	if len(pkgs) != 1 {
		return errors.New("more than one package found in files to be bundled")
	}
	pkg := pkgs["main"]

	// Need to move all import statements so they're at the front of
	// the new file; first, remove all existing import declarations.
	for _, f := range pkg.Files {
		decls := f.Decls[:0]
		for _, d := range f.Decls {
			if d, ok := d.(*ast.GenDecl); ok && d.Tok == token.IMPORT {
				continue
			}
			decls = append(decls, d)
		}
		f.Decls = decls
	}

	merged := ast.MergePackageFiles(pkg, ast.FilterUnassociatedComments|ast.FilterImportDuplicates)

	// Put imports back, using the parsed information.
	dec := &ast.GenDecl{
		Tok:    token.IMPORT,
		Lparen: 1, // kludge so ast/format outputs more than one import
	}
	for _, f := range pkg.Files {
		for _, i := range f.Imports {
			dec.Specs = append(dec.Specs, i)
		}
	}
	merged.Decls = append([]ast.Decl{dec}, merged.Decls...)

	if err := format.Node(tmp, fset, merged); err != nil {
		return err
	}

	if err := os.Rename(tmp.Name(), "asset.go"); err != nil {
		return err
	}
	tmp = nil
	return nil
}
