package main

import (
	. "fmt"
	"time"
)

func main() {
	arrayDefine()

	sliceDefine()

	mapDefine()

	chanDefine()

	structDefine()
}

// 数组初始化时长度已经定义，不能修改
func arrayDefine() {
	var i1 [1]int
	var i2 [1]*int
	i3 := [1]int{0}
	i4 := [2]int{1: 1}
	i5 := [...]int{1: 1}

	// not defined -> panic: runtime error: index out of range
	// i1[0] = 0

	// cannot use [2]int literal (type [2]int) as type [1]int in assignment
	// i1 = [2]int{}

	i1 = [1]int{}
	i1[0] = 0

	i2 = [1]*int{0: new(int)}
	*i2[0] = 0

	i3[0] = 0

	Println("array:", i1, i2, *i2[0], i3, i4, i5)
}

func sliceDefine() {
	sourceI := [5]int{0, 1, 2, 3, 4}

	s1 := sourceI[:]

	s2 := sourceI[1:]

	s3 := sourceI[2:3]

	s4 := sourceI[:len(sourceI)]

	// invalid slice index 100 (out of bounds for 5-element array)
	// s5 := sourceI[:100]

	Println("slice:", s1, s2, s3, s4)
}

func mapDefine() {
	var m1 map[string]int
	m2 := map[string]int{}
	m3 := make(map[string]int)

	// not defined  ->  panic: assignment to entry in nil map
	// m1["haah"] = 0

	m1 = make(map[string]int)
	m1["a"]++
	m1["b"] = 0

	m2["a"]++
	m2["b"] = 0

	m3["a"]++
	m3["b"] = 0

	Println("map:", m1, m2, m3)
}

// channel two-way binding
func chanDefine() {
	var c1 chan int
	var c2 chan int = make(chan int)
	var c3 chan int = make(chan int, 2)

	// all goroutines are asleep - deadlock!
	// c1 <- 1

	// all goroutines are asleep - deadlock!
	// c1 = make(chan int)
	// c1 <- 1
	// v1 := <-c1

	// init_complex_sample.go
	// groupChan(c1, 1, 2)

	go groupChan(c2, 1, 2)
	Println("sleep 1 second, wait for goroutine to finished")
	time.Sleep(1 * time.Second)

	go groupChan(c3, 1, 2, 3)
	Println("sleep 1 second, wait for goroutine to finished")
	time.Sleep(1 * time.Second)

	c4 := make(chan int, 1)
	go groupChanBulk(c4, 1, 2, 3, 4)
	Println("sleep 1 second, wait for goroutine to finished")
	time.Sleep(1 * time.Second)

	Println("chan:", c1, c2, c3)
	time.Sleep(1 * time.Second)
}

func groupChan(c chan int, vs ...int) {
	index := new(int)
	for _, v := range vs {
		go cChan(c, v)
		*index++
	}

	for {
		select {
		case v, ok := <-c:
			if !ok {
				return
			}
			Printf("value[%d]:%d", *index, v)
			Println()
		}
	}
}

func cChan(c chan int, v int) {
	c <- v
}

func cChanBulk(c chan int, vs []int) {
	for _, v := range vs {
		c <- v
	}
	Printf("bulk finished, value = %v, length = %d", vs, len(c))
	Println()
}

func groupChanBulk(c chan int, vs ...int) {
	go cChanBulk(c, vs)

	for {
		select {
		case v, ok := <-c:
			if !ok {
				return
			}
			Printf("value:%d", v)
			Println()
		}
	}
}

type testStruct struct {
	ppap int
	papp string
}

func structDefine() {
	s1 := new(testStruct)
	var s2 testStruct
	s3 := testStruct{papp: "qqq"}
	s4 := testStruct{321, "wahaha"}

	s1.ppap = 1
	Println(s1)

	s2.ppap = 2
	Println(s2)

	Println(s3)

	Println(s4)
}
