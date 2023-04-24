package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func main() {
	n := 12
	fmt.Println(numSquares(n))
}
