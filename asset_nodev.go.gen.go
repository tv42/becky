// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package main

var asset_nodev = asset(asset.init(asset{Name: "asset_nodev.go", Content: "" +
	"// +build !dev\n\npackage main\n\nimport (\n\t\"encoding/base64\"\n\t\"hash/fnv\"\n\t\"io\"\n\t\"net/http\"\n\t\"strings\"\n\t\"time\"\n)\n\ntype asset struct {\n\tName    string\n\tContent string\n\tetag    string\n}\n\nfunc (a asset) init() asset {\n\t// This is a method to minize the namespace pollution. Use\n\t// chaining to make it callable in variable declarations.\n\th := fnv.New64a()\n\t_, _ = io.WriteString(h, a.Content)\n\ta.etag = `\"` + base64.StdEncoding.EncodeToString(h.Sum(nil)) + `\"`\n\treturn a\n}\n\nfunc (a asset) ServeHTTP(w http.ResponseWriter, req *http.Request) {\n\tif a.etag != \"\" && w.Header().Get(\"ETag\") == \"\" {\n\t\tw.Header().Set(\"ETag\", a.etag)\n\t}\n\tbody := strings.NewReader(a.Content)\n\thttp.ServeContent(w, req, a.Name, time.Time{}, body)\n}\n" +
	""}))
