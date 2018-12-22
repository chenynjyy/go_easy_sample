package main

import (
	. "fmt"
	"math/rand"
	"time"
)

func main() {
	// c := make(chan<- int)
	// c := make(<-chan int)
	c := make(chan int)

	rand.Seed(3)

	go receive(c)
	go send(c)

	Println("Running by themselves")
	time.Sleep(10 * time.Second)
	Println("Run finished")
}

// restrict chan use on send
// 限制发送
func receive(c chan<- int) {
	for {
		c <- rand.Int()
		time.Sleep(1 * time.Second)
	}
}

// restrict chan use on receive
// 限制接收
func send(c <-chan int) {
	for {
		select {
		case v := <-c:
			Println(v)
		}
	}
}
