  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ cp -- "$TESTDIR/cat.go" .
  $ cp -- "$TESTDIR/../asset.go" .
  $ echo Hello, world >greeting.txt
  $ go generate
  $ go build -o cat
  $ ./cat
  Hello, world
