package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 0 {
		n := len(nums) + 1
		target := sum / 2
		dp := make([]bool, target+1)
		dp[0] = true
		for i := 1; i < n; i++ {
			for j := target; j >= nums[i-1]; j-- {
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
		return dp[target]
	}
	return false
}
func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
