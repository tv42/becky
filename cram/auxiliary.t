  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject

Using `-lib=false` disables creating `asset_*.gen.go` files:

  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ becky -lib=false greeting.txt
  $ find -type f -printf '%P\n' | sort
  go.mod
  greeting.txt
  greeting.txt.gen.go
  main.go

Leaving it out gets them created:

  $ becky greeting.txt
  $ find -type f -printf '%P\n' | sort
  asset_dev.gen.go
  asset_nodev.gen.go
  go.mod
  greeting.txt
  greeting.txt.gen.go
  main.go
