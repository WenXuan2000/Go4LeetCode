package main

import "fmt"

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return nums[low]
}

func main() {
	nums := []int{3, 0, 1, 1, 3}
	fmt.Println(findMin(nums))

}
