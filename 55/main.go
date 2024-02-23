package main

import "fmt"

func canJump(nums []int) bool {
	flag := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= flag-i {
			flag = i
		}
	}
	if flag == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
