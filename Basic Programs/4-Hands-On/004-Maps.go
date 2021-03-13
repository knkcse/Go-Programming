package main

import "fmt"

func main() {
	mp := map[string][]string{
		"Naveen": []string{"Kumar Kammari", "Loves coding", "Play cricket", "Shadow"},
		"Anil":   []string{"Kammari", "Amazing guy", "CTO", "A Techie", "Good Person", "Learner"},
		"Ammu":   []string{"Pretty", "Cool"},
	}
	fmt.Println(mp)
	for k, x := range mp {
		fmt.Println("Key is ", k)
		for i, v := range x {
			fmt.Printf("Index: %v\tValue: %v\n", i, v)
		}
		fmt.Println()
	}

	mp["Varun"]=[]string{"Nephew","Good boy","Never cry","Happy face"}
	fmt.Println(mp)
	delete(mp,"Naveen")
	fmt.Println(mp)
}
