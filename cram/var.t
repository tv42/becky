  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ cp -- "$TESTDIR/cat-xyzzy.go" .
  $ echo Hello, world >greeting.txt
  $ becky -var=xyzzy greeting.txt
  $ go build -o cat-xyzzy
  $ ./cat-xyzzy
  Hello, world
