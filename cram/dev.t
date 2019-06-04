  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ cp -- "$TESTDIR/fakehttp.go" .
  $ cp -- "$TESTDIR/../asset.go" .
  $ echo Hello, world >greeting.txt
  $ go generate
  $ go build -tags=dev -o fakehttp
  $ ./fakehttp
  Hello, world
  $ rm greeting.txt
  $ ./fakehttp
  fakehttp.go: bad response: 500: open greeting.txt: no such file or directory
  [1]
