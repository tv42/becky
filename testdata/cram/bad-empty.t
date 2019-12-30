  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/../.."

  $ echo Hello, world >greeting.txt
  $ go run github.com/tv42/becky greeting.txt
  becky: cannot load package: go [list -e -json -compiled=false -test=false -export=false -deps=false -find=true -- .]: exit status 1: build .: cannot find module for path .
  exit status 1
  [1]
  $ find -type f -printf '%P\n' | sort
  go.mod
  go.sum
  greeting.txt
