  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >no-extension
  $ becky -var=xyzzy no-extension
  becky: files without extension need -wrap: no-extension
  [1]
