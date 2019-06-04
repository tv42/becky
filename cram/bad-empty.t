  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ echo Hello, world >greeting.txt
  $ becky greeting.txt
  becky: no buildable Go source files in .* (re)
  [1]
  $ find -type f -printf '%P\n' | sort
  go.mod
  greeting.txt
