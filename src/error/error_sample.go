package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	i1, e1 := errofF(true)
	i2, e2 := errofF(false)
	log.Println("function 1 :", i1, e1)
	log.Println("function 2 :", i2, e2)

	e3 := err()
	log.Println("error 3:", e3)
}

func errofF(ok bool) (int, error) {
	return 0, fmt.Errorf("something wrong")
}

func err() error {
	return errors.New("error happen")
}
