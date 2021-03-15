package main

import (
	"fmt"
	"sort"
)

type student struct {
	fname, lname string
	age          int
	id           string
}
type custom []student // custom implements sort.Interface based on the Age field.

func (s custom) Len() int {
	return len(s)
}
func (s custom) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s custom) Less(i, j int) bool {
	return s[i].age < s[j].age
}
func main() {
	fmt.Println("Custom Sort")
	s1 := student{
		"Naveen Kumar",
		"Kammari",
		22,
		"CS19MTECH11009",
	}
	s2 := student{
		"Anil Kumar",
		"Vadla",
		26,
		"B082597",
	}
	fmt.Println(s1)
	fmt.Println(s2)
	var studentList []student
	studentList = append(studentList, s1, s2)
	sort.Sort(custom(studentList)) //studentList will become custom which is of type Interface
	fmt.Println(studentList)

}
