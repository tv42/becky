package assets

//go:generate go run github.com/tv42/becky -var=Greeting greeting.txt

func txt(a asset) asset {
	return a
}
