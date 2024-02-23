package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	//index, k := 0, 1
	index := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			index += 1
		} else {
			index = 0
		}
		//if index == 0 {
		//	temp := nums[i:]
		//	copy(nums[k:], temp)
		//	n -= i - k
		//	i = k
		//	index = 0
		//	continue
		//}
		if index > 1 {
			temp := nums[i+1:]
			copy(nums[i:], temp)
			n--
			i--
		}
	}
	return n
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums), nums)
}
