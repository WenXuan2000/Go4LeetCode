package main

import "fmt"

//122. 买卖股票的最佳时机 II
//https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china

func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}
func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1,2,3,4,5}
	fmt.Println(maxProfit(prices))

}
