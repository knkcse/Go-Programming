package main

import "fmt"

type student struct {
	first, last, father string
	mobile              string
}
type record struct {
	student
	rcrd int
}

func main() {
	s1 := student{"Naveen Kumar", "Kammari", "K Ponnaji", "9491841838"}
	fmt.Println(s1)
	recordList := record{}
	recordList.student = s1
	recordList.rcrd = 14
	fmt.Println(recordList)
	//s2 := student{}
	//fmt.Scan(&s2.first, &s2.last, &s2.father, &s2.mobile)
	s2 := record{}
	fmt.Scan(&s2.first, &s2.last, &s2.father, &s2.mobile)
	fmt.Println(s2)
	fmt.Println(s2.first)
	fmt.Println(s2.student.first)

	//Array of Structs
	studentList := [2]record{}
	studentList[0].student = s1
	studentList[0].rcrd = 1
	studentList[1].student = s2.student
	studentList[1].rcrd = 2
	fmt.Println()
	fmt.Println(studentList)
}
