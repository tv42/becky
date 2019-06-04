  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ cp -- "$TESTDIR/cat.go" .
  $ cp -- "$TESTDIR/../asset.go" .
  $ echo Hello, world >greeting.txt
  $ go generate
  $ go build -o cat
  $ ./cat
  Hello, world
