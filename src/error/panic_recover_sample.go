package main

import "log"

func main() {
	test()
	println("test() is running well")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work fail:", err)
		}
	}()

	panic("something happen WTF!!!!!!")

	// compile success, but this function will never used
	println("heiheihei")
}
