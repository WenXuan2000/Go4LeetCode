package main

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = Min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
			if ans < dp[i][j]*dp[i][j] {
				ans = dp[i][j] * dp[i][j]
			}
		}
	}
	return ans
}
func Min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else if z <= x && z <= y {
		return z
	}
	return -1
}
func main() {
	//matrix := [][]byte{
	//	{'1', '0', '1', '0', '0'},
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '1', '1', '1', '1'},
	//	{'1', '0', '0', '1', '0'}}
	matrix := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'0', '0', '1', '1', '1'}}
	fmt.Println(maximalSquare(matrix))
}
