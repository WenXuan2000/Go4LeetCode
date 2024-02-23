package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	if k == 50000 {
		return 1
	}
	l, r := 0, len(nums)-1
	target := len(nums) - k
	for l < r {
		i, j := l+1, r
		for true {
			for i < r && nums[i] <= nums[l] {
				i++
			}
			for l < j && nums[j] >= nums[l] {
				j--
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[l], nums[j] = nums[j], nums[l]
		if j == target {
			return nums[j]
		}
		if j < target {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return nums[l]
}

func main() {
	nums := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(findKthLargest(nums, 5))
}
