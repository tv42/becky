  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ cp -- "$TESTDIR/fakehttp.go" .
  $ cp -- "$TESTDIR/../asset.go" .
  $ echo Hello, world >greeting.txt
  $ go generate
  $ go build -tags=dev -o fakehttp
  $ ./fakehttp
  Hello, world
  $ rm greeting.txt
  $ ./fakehttp
  fakehttp.go: bad response: 500: open (.*)/src/example.com/myproject/greeting.txt: no such file or directory (re)
  [1]
