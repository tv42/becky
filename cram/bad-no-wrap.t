  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >no-extension
  $ becky -var=xyzzy no-extension
  becky: files without extension need -wrap: no-extension
  [1]
