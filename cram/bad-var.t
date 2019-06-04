  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ becky -var=not-a-go-identifier greeting.txt
  becky: not a valid go identifier: not-a-go-identifier
  [1]
