package main

// function and method are not the same
func main() {
	var ii iiInterface = new(testStruct)
	println(ii.test())

	println(test())

	var iii iiiInterface = new(testIStruct)
	println(iii.test())
	println(iii.gg())

}

type iiInterface interface {
	test() int
}

type iiiInterface interface {
	iiInterface
	gg() int
}

type testStruct struct {
	element1 int
	element2 string
}

type testIStruct struct {
	testStruct
	element3 string
}

func (s *testIStruct) test() (i int) {
	i = 333
	return
}

func (s *testIStruct) gg() int {
	s.element1 = 1
	return 321
}

func (s testIStruct) gg() int {
	s.element1 = 1
	return 321
}

// this is method
func (st *testStruct) test() (i int) {
	i = 123
	return
}

// this is function
func test() (i int) {
	i = 123
	return
}
