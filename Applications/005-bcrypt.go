package main

//package bcrypt
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pswd := "rgukt123"
	bs := []byte(pswd)
	fmt.Println(bs)
	hash, err := bcrypt.GenerateFromPassword(bs, bcrypt.MinCost) //kind of encrypting password
	fmt.Println(hash, err)

	error := bcrypt.CompareHashAndPassword(hash, []byte("rguddkt123"))
	if error != nil {
		fmt.Println("Your password is wrong")
	} else {
		fmt.Println("Your password is correct")
	}

}
