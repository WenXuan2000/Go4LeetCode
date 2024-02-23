package main

import "fmt"

func longestAlternatingSubarray(nums []int, threshold int) int {
	var temp int
	ans := 0
	if nums[0]%2 == 0 && nums[0] <= threshold {
		temp = 1
		ans = 1
	} else {
		temp = 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]%2 == 0 && temp == 0 && nums[i] <= threshold {
			temp += 1
		} else if temp != 0 && nums[i] <= threshold {
			if nums[i]%2 != nums[i-1]%2 {
				temp += 1
			} else if nums[i]%2 == 0 {
				temp = 1
			} else if nums[i]%2 == 1 {
				temp = 0
			}
		} else {
			temp = 0
		}
		if temp > ans {
			ans = temp
		}
	}
	return ans
}

func main() {
	nums := []int{2, 10, 5}
	threshold := 7
	fmt.Println(longestAlternatingSubarray(nums, threshold))
}
