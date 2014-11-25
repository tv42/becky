  $ export GOPATH="$PWD"
  $ mkdir -p src/example.com/myproject
  $ cd src/example.com/myproject
  $ echo Hello, world >greeting.txt
  $ becky greeting.txt
  becky: no buildable Go source files in .* (re)
  [1]
  $ find -type f -printf '%P\n' | sort
  greeting.txt
