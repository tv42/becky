  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ echo package xyzzy >xyzzy.go
  $ echo Hello, world >greeting.txt
  $ becky greeting.txt
  $ grep ^package *.gen.go | sort
  asset_dev.gen.go:package xyzzy
  asset_nodev.gen.go:package xyzzy
  greeting.txt.gen.go:package xyzzy
