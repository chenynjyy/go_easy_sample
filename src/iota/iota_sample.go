package main

const (
	_ = iota
	B
	C
	D = 10 * iota
	E = 1 << (iota * 5)
)

func main() {

	println(B, C, D, E)
}
