  $ go mod init example.com/myproject
  go: creating new go.mod: module example.com/myproject
  $ go mod edit -replace=github.com/tv42/becky="$TESTDIR/../.."

  $ go run github.com/tv42/becky
  Usage:
    becky [OPTS] FILE..
  
  Creates files FILE.gen.go and asset_*.gen.go
  
  Options:
    -lib
      \tgenerate asset_*.gen.go files defining the asset type (default true) (esc)
    -var string
      \tvariable name to use, "_" to ignore (default: file basename without extension) (esc)
    -wrap string
      \twrapper function or type (default: filename extension) (esc)
  exit status 2
  [1]

  $ go run github.com/tv42/becky -help
  Usage:
    becky [OPTS] FILE..
  
  Creates files FILE.gen.go and asset_*.gen.go
  
  Options:
    -lib
      \tgenerate asset_*.gen.go files defining the asset type (default true) (esc)
    -var string
      \tvariable name to use, "_" to ignore (default: file basename without extension) (esc)
    -wrap string
      \twrapper function or type (default: filename extension) (esc)
  exit status 2
  [1]
