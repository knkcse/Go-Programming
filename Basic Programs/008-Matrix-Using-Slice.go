package main

import "fmt"

func main() {
	fmt.Println("Enter your matrix size of the form MxN")
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)
	var matrix = [][]int{} //For 2-D matrix
	for i := 0; i < m; i++ {
		var row = []int{}
		var x int
		for j := 0; j < n; j++ {
			fmt.Scan(&x)
			row = append(row, x)
		}
		matrix = append(matrix, row)
	}
	fmt.Println()
	fmt.Println(matrix)
	fmt.Println("Finding the maximum and minimum elements in the matrix")
	var maxE, minE int = matrix[0][0], matrix[0][0]
	fmt.Println()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j],"  ")
			if maxE < matrix[i][j] {
				maxE = matrix[i][j]
			}
			if minE > matrix[i][j] {
				minE = matrix[i][j]
			}

		}
		fmt.Println()
	}
	fmt.Println("Maximum Element: ", maxE)
	fmt.Println("Minimum Element: ", minE)
	fmt.Println()
	for i,x:= range matrix{
		for j,v:=range x{
			fmt.Println("Row: ",i," Column: ",j," Elemen: ",v)
		}
	}

}
