package main

import "fmt"

func main() {

	var v1 int
	var v2 bool
	var v3 string
	var v4 map[string]int
	var v5 interface{}
	var v6 []interface{}
	var v7 float32

	fmt.Printf("int = [%v]", v1)
	fmt.Println()
	fmt.Printf("bool = [%v]", v2)
	fmt.Println()
	fmt.Printf("string = [%v]", v3)
	fmt.Println()
	fmt.Printf("map = [%v]", v4)
	fmt.Println()
	fmt.Printf("interface = [%v]", v5)
	fmt.Println()
	fmt.Printf("[]interface = [%v]", v6)
	fmt.Println()
	fmt.Printf("float32 = [%v]", v7)
	fmt.Println()

}
