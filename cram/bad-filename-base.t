Empty file basenames cannot be used as variable name.

  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >.txt
  $ becky .txt
  becky: cannot use empty file basename as identifier: .txt
  [1]
