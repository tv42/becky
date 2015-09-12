// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package main

var assetDev = asset(asset{Name: "asset_dev.go", Content: "" +
	"// +build dev\n\npackage main\n\nimport (\n\t\"go/build\"\n\t\"net/http\"\n\t\"os\"\n\t\"path/filepath\"\n\t\"time\"\n)\n\ntype asset struct {\n\tName    string\n\tContent string\n\t// don't bother precomputing ETag if we're reloading from disk\n}\n\nfunc (a asset) init() asset {\n\treturn a\n}\n\nfunc (a asset) importPath() string {\n\t// filled at code gen time\n\treturn \"{{.ImportPath}}\"\n}\n\nfunc (a asset) Open() (*os.File, error) {\n\tpath := a.importPath()\n\tpkg, err := build.Import(path, \".\", build.FindOnly)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tp := filepath.Join(pkg.Dir, a.Name)\n\treturn os.Open(p)\n}\n\nfunc (a asset) ServeHTTP(w http.ResponseWriter, req *http.Request) {\n\tbody, err := a.Open()\n\tif err != nil {\n\t\t// show the os.Open message, with paths and all, but this only\n\t\t// happens in dev mode.\n\t\thttp.Error(w, err.Error(), http.StatusInternalServerError)\n\t\treturn\n\t}\n\tdefer body.Close()\n\thttp.ServeContent(w, req, a.Name, time.Time{}, body)\n}\n" +
	"", etag: `"c1Es6vaoVm4="`})
