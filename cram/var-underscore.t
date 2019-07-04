  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/.."

  $ echo package main >main.go
  $ echo foo >one.txt
  $ echo bar >two.txt
  $ go run github.com/tv42/becky -var=_ one.txt two.txt
  $ grep ^var one.txt.gen.go two.txt.gen.go | cut -d= -f1
  one.txt.gen.go:var _ 
  two.txt.gen.go:var _ 
