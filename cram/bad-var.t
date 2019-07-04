  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/.."

  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ go run github.com/tv42/becky -var=not-a-go-identifier greeting.txt
  becky: not a valid go identifier: not-a-go-identifier
  exit status 1
  [1]
