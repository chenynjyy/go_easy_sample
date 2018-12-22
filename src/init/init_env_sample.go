package main

import "os"

var (
	v1 = os.Getenv("GOPATH")
)

func main() {
	println(v1)
}
