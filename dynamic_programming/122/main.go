package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	if len(prices) < 2 {
		return 0
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
