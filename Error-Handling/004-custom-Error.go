package main

import "fmt"

type customError struct {
	arg int
	err string
}

//We need to  attached with Error() method
func (c customError) Error() string {
	return fmt.Sprintf("Error is - %v %v", c.arg, c.err)
}
func main() {
	err:=foo(3)
	if err != nil {
		fmt.Println("Error is: ",err)
		return
	}
	fmt.Println("No error")
}
func foo(a int) error {
	//fmt.Println(c)
	if a<18{
		return customError{a,"He is not eligible"}
	}
	return nil
}

