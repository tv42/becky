  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/../.."

  $ echo package xyzzy >xyzzy.go
  $ echo Hello, world >greeting.txt
  $ go run github.com/tv42/becky greeting.txt
  $ grep ^package *.gen.go | sort
  asset_dev.gen.go:package xyzzy
  asset_nodev.gen.go:package xyzzy
  greeting.txt.gen.go:package xyzzy
