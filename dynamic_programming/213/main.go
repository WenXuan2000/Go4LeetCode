package main

import (
	"fmt"
)

func _rob(nums []int) int {
	first, second := nums[0], Max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, Max(first+v, second)
	}
	return second
}
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return Max(nums[0], nums[1])
	}
	return Max(_rob(nums[:n-1]), _rob(nums[1:]))
}
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}
