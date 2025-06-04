package main

import "fmt"

func main() {

	matrix1 := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	matrix2 := [][]int{{1, 2}, {2, 2}}
	matrix3 := [][]int{{11, 74, 0, 93}, {40, 11, 74, 7}}

	fmt.Println(isToeplitzMatrix(matrix1))
	fmt.Println(isToeplitzMatrix(matrix2))
	fmt.Println(isToeplitzMatrix(matrix3))

}

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
