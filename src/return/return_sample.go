package main

import . "fmt"

func main() {
	println(single())
	v1, v2 := multi()
	Printf("%d, %s", v1, v2)

	// multiple-value multi() in single-value context
	// Println(multi())

	v3, _, v4, _ := multiNil()
	Println(v3, v4)
}

func single() int {
	return 123
}

func multi() (int, string) {
	return 321, "wahaha"
}

func multiNil() (int, string, string, map[string]int) {
	return 123, "", "gg", nil
}
