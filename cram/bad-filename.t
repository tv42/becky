Files with basenames that are not Go identifiers are errors, unless
the `-var` flag is used.

  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ cp -- "$TESTDIR/cat-xyzzy.go" .
  $ echo Hello, world >not-a-go-identifier.txt
  $ becky not-a-go-identifier.txt
  becky: not a valid go identifier: not-a-go-identifier
  [1]
  $ becky -var=xyzzy not-a-go-identifier.txt
  $ go build -o cat-xyzzy
  $ ./cat-xyzzy
  Hello, world
