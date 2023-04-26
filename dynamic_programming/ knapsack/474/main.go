package main

import (
	"fmt"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros, ones := Count(s)
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = Max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Count(s string) (int, int) {
	zeros := strings.Count(s, "0")
	return zeros, len(s) - zeros
}
func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
}
