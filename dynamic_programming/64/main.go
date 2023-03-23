package main

import "fmt"

//func minPathSum(grid [][]int) int {
//	ans := 0
//	m := len(grid)
//	n := len(grid[0])
//	if n == 0 {
//		return 0
//	} else if n == 1 {
//		for i := 0; i < m; i++ {
//			ans += grid[i][0]
//		}
//		return ans
//	}
//	if m == 1 {
//		for i := 0; i < n; i++ {
//			ans += grid[0][i]
//		}
//		return ans
//	}
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, n)
//	}
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if i == 0 && j == 0 {
//				dp[0][0] = grid[0][0]
//			} else if i == 0 {
//				dp[0][j] = dp[0][j-1] + grid[0][j]
//			} else if j == 0 {
//				dp[i][0] = dp[i-1][0] + grid[i][0]
//			} else {
//				dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
//			}
//		}
//	}
//	return dp[m-1][n-1]
//}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[0] = grid[0][0]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[0][j]
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][0]
			} else {
				dp[j] = Min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

/*
1 3 1
1 5 1
4 2 1
*/
func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
