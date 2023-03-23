package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	length := len(nums)
	pre2, pr1 := 0, 0
	var cur int
	for i := 0; i < length; i++ {
		cur = Max(pre2+nums[i], pr1)
		pre2 = pr1
		pr1 = cur
	}
	return cur
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
