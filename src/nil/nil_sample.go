package main

import (
	. "fmt"
	// "reflect"
)

func main() {
	var stringV string
	var intV int
	var boolV bool
	var int64V int64
	var mapV map[string]int
	var chanV chan int

	print(stringV, intV, boolV, int64V, mapV, chanV)
}

func print(values ...interface{}) {
	for index, value := range values {
		// typeT := reflect.TypeOf(value)
		// Println(typeT.Align(), typeT.FieldAlign(), typeT.NumMethod(), typeT.Name(), typeT.PkgPath(), typeT.PkgPath(), typeT.Size(), typeT.String(), typeT.Kind())
		Printf("index = [%d], value type = [%T], value = [%v]", index, value, value)
		Println()
	}
}
