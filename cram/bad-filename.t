  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/.."

Files with basenames that are not Go identifiers are errors, unless
the `-var` flag is used.

  $ cp -- "$TESTDIR/cat-xyzzy.go" .
  $ echo Hello, world >not-a-go-identifier.txt
  $ go run github.com/tv42/becky not-a-go-identifier.txt
  becky: not a valid go identifier: not-a-go-identifier
  exit status 1
  [1]
  $ go run github.com/tv42/becky -var=xyzzy not-a-go-identifier.txt
  $ go build -o cat-xyzzy
  $ ./cat-xyzzy
  Hello, world
