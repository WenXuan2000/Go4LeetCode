package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}
func main() {
	text1, text2 := "abcde", "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
