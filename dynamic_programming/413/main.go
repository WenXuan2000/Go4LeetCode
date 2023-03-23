package main

import "fmt"

//func numberOfArithmeticSlices(nums []int) int {
//	if len(nums) < 3 {
//		return 0
//	}
//	dp := make([]int, len(nums))
//	for i := 2; i < len(nums); i++ {
//		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
//			dp[i] = dp[i-1] + 1
//		}
//	}
//	ans := 0
//	for _, v := range dp {
//		ans += v
//	}
//	return ans
//}
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	ans := 0
	index := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			index += 1
			ans += index
		} else {
			index = 0
		}
	}
	return ans
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(numberOfArithmeticSlices(nums))
}
