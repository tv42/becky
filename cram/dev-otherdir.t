  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/.."

  $ cp -- "$TESTDIR/fakehttp-otherdir.go" .
  $ mkdir assets
  $ echo Hello, world >assets/greeting.txt
  $ cp -- "$TESTDIR/fakehttp-otherdir-assets.go" assets/assets.go
  $ go generate ./assets
  $ go build -tags=dev -o fakehttp
  $ ./fakehttp
  Hello, world
  $ rm assets/greeting.txt
  $ ./fakehttp
  fakehttp.go: bad response: 500: open .*/assets/greeting.txt: no such file or directory (re)
  [1]
