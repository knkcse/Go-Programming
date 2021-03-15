package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	s := `[{"First":"Naveen","Last":"Kammari","Age":22},{"First":"Anil","Last":"Kumar","Age":26}]`
	fmt.Println(s)
	bs := []byte(s)
	//fmt.Println(bs)
	fmt.Println()
	var people []person
	err := json.Unmarshal(bs, &people) //Unmarshaling json data back to slice people of person type
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println(people)
	}

	//Printing
	for i, v := range people {
		fmt.Println("Person: ", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
	}

}
