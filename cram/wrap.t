  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ becky -wrap=xyzzy greeting.txt
  $ grep ^var greeting.txt.gen.go | cut -d'(' -f1
  var greeting = xyzzy
