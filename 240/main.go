package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] > target {
		return false
	}
	n := len(matrix[0]) - 1
	for i, j := 0, n; i < len(matrix) && j >= 0; {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else if matrix[i][j] == target {
			return true
		}
	}
	return false
}
func main() {
	//eval_input.Eval_input("[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]")
	matrix := [][]int{{-1, 3}}
	target := 5
	fmt.Println(searchMatrix(matrix, target))
}
