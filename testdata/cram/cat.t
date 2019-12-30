  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/../.."

  $ cp -- "$TESTDIR/cat.go" .
  $ echo Hello, world >greeting.txt
  $ go generate
  $ go build -o cat
  $ ./cat
  Hello, world
