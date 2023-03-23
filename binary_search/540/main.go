package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1
	if len(nums) == 1 {
		return nums[0]
	}
	for left < right {
		mid := (left + right) / 2
		if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
			return nums[mid]
		}
		//前面是偶数
		if (mid-left)%2 == 0 {
			if nums[mid] == nums[mid+1] {
				left = mid + 2
			} else if nums[mid] == nums[mid-1] {
				right = mid - 2
			}
			//	前面是奇数
		} else {
			if nums[mid] == nums[mid+1] {
				right = mid - 1
			} else if nums[mid] == nums[mid-1] {
				left = mid + 1
			}
		}
	}
	return nums[left]
}

func main() {
	//nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	//nums := []int{3, 3, 7, 7, 10, 11, 11}
	nums := []int{1, 1, 2, 3, 3}
	fmt.Println(singleNonDuplicate(nums))
}
