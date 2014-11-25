  $ SRC="$(go list -f '{{.Dir}}' github.com/tv42/becky)"
  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject

Using `-lib=false` disables creating `asset_*.gen.go` files:

  $ echo package main >main.go
  $ echo Hello, world >greeting.txt
  $ becky -lib=false greeting.txt
  $ find -type f -printf '%P\n' | sort
  greeting.txt
  greeting.txt.gen.go
  main.go

Leaving it out gets them created:

  $ becky greeting.txt
  $ find -type f -printf '%P\n' | sort
  asset_dev.gen.go
  asset_nodev.gen.go
  greeting.txt
  greeting.txt.gen.go
  main.go
