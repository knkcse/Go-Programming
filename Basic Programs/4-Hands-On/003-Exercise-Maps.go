package main

import "fmt"

func main() {
	mp := map[string]string{
		"Naveen": "Love to learn new things",
		"Anil":   "He's an amazing guy!!",
		"Kavi":   "She is their sister",
	}
	//mp["Naveen"] = "Love to learn new things"
	//mp["Anil"] = "He's an amazing guy!!"
	//mp["Kavi"] = "She is their sister"
	fmt.Println(mp)
	for k,v:=range mp{
		fmt.Println(k,v)
	}
}
