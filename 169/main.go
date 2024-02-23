package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	dict := make(map[int]int)
	ans := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] += 1
		if dict[nums[i]] > temp {
			temp = dict[nums[i]]
			ans = nums[i]
		}
		if dict[nums[i]] > int(len(nums)/2) {
			break
		}
	}
	return ans
}

//func majorityElement(nums []int) int {
//
//}
func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
