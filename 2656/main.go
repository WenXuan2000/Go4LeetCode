package main

func maximizeSum(nums []int, k int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	ans := max*k + ((k-1)*k)/2
	return ans
}
