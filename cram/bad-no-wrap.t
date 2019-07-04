  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/.."

  $ echo package main >main.go
  $ echo Hello, world >no-extension
  $ go run github.com/tv42/becky -var=xyzzy no-extension
  becky: files without extension need -wrap: no-extension
  exit status 1
  [1]
