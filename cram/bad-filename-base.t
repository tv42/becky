Empty file basenames cannot be used as variable name.

  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ echo package main >main.go
  $ echo Hello, world >.txt
  $ becky .txt
  becky: cannot use empty file basename as identifier: .txt
  [1]
