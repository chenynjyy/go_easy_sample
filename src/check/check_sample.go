package main

import . "fmt"

func main() {
	var tester Tester = new(TestStruct)
	if testStruct, ok := tester.(*TestStruct); ok {
		Println("type:", testStruct)
		testStruct.test()
		Println(testStruct)
	} else {
		Println("type:", testStruct)
	}
}

type Tester interface {
	test()
}

type TestStruct struct {
	no     int
	dragon string
	age    uint
}

func (t *TestStruct) test() {
	t.age = 1
	t.dragon = "wangnima"
	t.age = 1
}
