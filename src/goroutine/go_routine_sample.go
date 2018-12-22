package main

import "time"

func main() {
	go test()

	// if this function not present, goroutine test() may not execute finish
	time.Sleep(1 * time.Second)
}

func test() {
	println("test")
}
