package main

import "fmt"

func quick_sort(nums []int, l, r int) []int {
	if l < r {
		pre := l
		index := l + 1
		for i := index; i <= r; i++ {
			if nums[pre] > nums[i] {
				nums[index], nums[i] = nums[i], nums[index]
				index++
			}
		}
		nums[pre], nums[index-1] = nums[index-1], nums[pre]
		quick_sort(nums, l, index-1)
		quick_sort(nums, index, r)
	}
	return nums
}

func merge_sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	i := len(nums) / 2
	left := merge_sort(nums[:i])
	right := merge_sort(nums[i:])
	result := merge(left, right)
	return result
}
func merge(left []int, right []int) []int {
	i, j := 0, 0
	l, r := len(left), len(right)
	results := make([]int, 0)
	for i < l && j < r {
		if left[i] < right[j] {
			results = append(results, left[i])
			i++
		} else {
			results = append(results, right[j])
			j++
		}
	}
	results = append(results, left[i:]...)
	results = append(results, right[j:]...)
	return results
}

func insertion_sort(nums []int, n int) []int {
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

func bubble_sort(nums []int, n int) []int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func selection_sort(nums []int, n int) []int {
	for i := 0; i < n; i++ {
		index := i

		for j := i + 1; j < n; j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
	return nums
}
func main() {
	nums := []int{1, 3, 5, 7, 2, 6, 4, 8, 9, 2, 8, 7, 6, 0, 3, 5, 9, 4, 1, 0}
	l, r := 0, len(nums)-1
	fmt.Println("before sort :", nums)
	fmt.Println("after quick_sort:", quick_sort(nums, l, r))
	fmt.Println("after merge_sort:", merge_sort(nums))
	fmt.Println("after insertion_sort:", insertion_sort(nums, r+1))
	fmt.Println("after bubble_sort:", bubble_sort(nums, r+1))
	fmt.Println("after selection_sort:", selection_sort(nums, r+1))
}
