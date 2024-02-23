package main

import "fmt"

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	copy(ans, prices)
	st := []int{}
	for i, x := range prices {
		for len(st) != 0 && x <= prices[st[len(st)-1]] {
			ans[st[len(st)-1]] -= x
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
func main() {
	prices := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(prices))
}
