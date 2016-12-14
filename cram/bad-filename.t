Files with basenames that are not Go identifiers are errors, unless
the `-var` flag is used.

  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ cp -- "$TESTDIR/cat-xyzzy.go" .
  $ echo Hello, world >not-a-go-identifier.txt
  $ becky not-a-go-identifier.txt
  becky: not a valid go identifier: not-a-go-identifier
  [1]
  $ becky -var=xyzzy not-a-go-identifier.txt
  $ go build -o cat-xyzzy
  $ ./cat-xyzzy
  Hello, world
