  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ becky -var=not-a-go-identifier greeting.txt
  becky: not a valid go identifier: not-a-go-identifier
  [1]
