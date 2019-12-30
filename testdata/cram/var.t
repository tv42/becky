  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/../.."

  $ cp -- "$TESTDIR/cat-xyzzy.go" .
  $ echo Hello, world >greeting.txt
  $ go run github.com/tv42/becky -var=xyzzy greeting.txt
  $ go build -o cat-xyzzy
  $ ./cat-xyzzy
  Hello, world
