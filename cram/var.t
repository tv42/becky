  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ cp -- "$TESTDIR/cat-xyzzy.go" .
  $ echo Hello, world >greeting.txt
  $ becky -var=xyzzy greeting.txt
  $ go build -o cat-xyzzy
  $ ./cat-xyzzy
  Hello, world
