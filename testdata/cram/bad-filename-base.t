  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/../.."

Empty file basenames cannot be used as variable name.

  $ echo package main >main.go
  $ echo Hello, world >.txt
  $ go run github.com/tv42/becky .txt
  becky: cannot use empty file basename as identifier: .txt
  exit status 1
  [1]
