  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ echo package main >main.go
  $ echo foo >one.txt
  $ echo bar >two.txt
  $ becky -var=_ one.txt two.txt
  $ grep ^var one.txt.gen.go two.txt.gen.go | cut -d= -f1
  one.txt.gen.go:var _ 
  two.txt.gen.go:var _ 
