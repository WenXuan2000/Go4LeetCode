package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				equalOp := 0
				if word1[i-1] == word2[j-1] {
					equalOp = dp[i-1][j-1] + 0
				} else {
					equalOp = dp[i-1][j-1] + 1
				}
				dp[i][j] = Min(equalOp, Min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}
	return dp[m][n]
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
