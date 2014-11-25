  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ becky -wrap=xyzzy greeting.txt
  $ grep ^var greeting.txt.gen.go | cut -d'(' -f1
  var greeting = xyzzy
