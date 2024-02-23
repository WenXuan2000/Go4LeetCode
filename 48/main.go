package main

import (
	"fmt"
	"go4leetcode/eval_input"
)

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	var temp int
	for i := 0; i <= n/2; i++ {
		for j := i; j < n-i; j++ {
			temp = matrix[j][n-i]
			matrix[j][n-i] = matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = temp
		}
	}

}

func main() {
	eval_input.Eval_input("[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]")
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	rotate(matrix)
	fmt.Println(matrix)
}
