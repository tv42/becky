  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/../.."

  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ go run github.com/tv42/becky -wrap=xyzzy greeting.txt
  $ grep ^var greeting.txt.gen.go | cut -d'(' -f1
  var greeting = xyzzy
