package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 2
	}
	dp[0] = 0
	for i := 1; i <= len(coins); i++ {
		for j := coins[i-1]; j <= amount; j++ {
			dp[j] = Min(dp[j], dp[j-coins[i-1]]+1)
		}
	}
	if dp[amount] == amount+2 {
		return -1
	}
	return dp[amount]
}
func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}

}
func main() {
	coins := []int{3}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
