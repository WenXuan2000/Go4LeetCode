package main

import "fmt"

func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	r := 1
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		l[i] = l[i] * r
		r *= nums[i]
	}
	return l
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
