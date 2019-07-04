  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/.."

  $ echo Hello, world >greeting.txt
  $ go run github.com/tv42/becky greeting.txt
  becky: no buildable Go source files in .* (re)
  exit status 1
  [1]
  $ find -type f -printf '%P\n' | sort
  go.mod
  greeting.txt
