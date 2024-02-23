package main

import (
	"fmt"
)

//func maxProfit(prices []int) int {
//	n := len(prices)
//	buy := 100000
//	profit := 0
//	for i := 0; i < n; i++ {
//		buy = Min(prices[i], buy)
//		profit = Max(profit, prices[i]-buy)
//	}
//	return profit
//}
func maxProfit(prices []int) int {
	ans := 0
	buy := 10000
	for _, x := range prices {
		buy = Min(buy, x)
		ans = Max(ans, x-buy)
	}
	return ans
}
func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
