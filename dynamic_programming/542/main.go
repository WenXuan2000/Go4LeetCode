package main

import (
	"fmt"
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if n == 1 && m == 1 {
		return mat
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32 - 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = Min(dp[i][j], dp[i][j-1]+1)
				}
				if i > 0 {
					dp[i][j] = Min(dp[i][j], dp[i-1][j]+1)
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if j < n-1 {
					dp[i][j] = Min(dp[i][j], dp[i][j+1]+1)
				}
				if i < m-1 {
					dp[i][j] = Min(dp[i][j], dp[i+1][j]+1)
				}
			}
		}
	}
	return dp
}

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	//[[0,0,0],[0,1,0],[1,1,1]]
	fmt.Println(updateMatrix(mat))
}
