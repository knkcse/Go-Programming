package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("This is basic example of erros")
	value, err := funcFOO(34)
	if err != nil {
		fmt.Println("Here is the error message: ", err)
		fmt.Printf("TYPE: %T\n", err)
		return
	}
	fmt.Println(value)
}
func funcFOO(a int) (int, error) { //error build-in type
	if a%2 == 0 {
		return -1, errors.New("We are expecting odd numbers")
	}
	return a + 2, nil
}
