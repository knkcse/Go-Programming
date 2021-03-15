package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type fruits struct {
	Page   int
	Fruits []string
}

func main() {
	fmt.Println("For json string use NewEncoder")
	s := fruits{
		1,
		[]string{
			"Naveen ",
			"Anil",
			"Kammari",
		},
	}
	//bs,err:=json.Marshal(s)
	fmt.Println(s)
	fmt.Println()
	//if err!=nil{
	//	fmt.Println("Error: ",err)
	//}

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(s)
	enc.Encode(d)
}
