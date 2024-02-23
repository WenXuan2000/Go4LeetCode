package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	var dfs func(int)
	dfs = func(level int) {
		if level == len(nums)-1 {
			temp := make([]int, 0)
			temp = append(temp, nums...)
			res = append(res, temp)
			return
		}
		for i := level; i < len(nums); i++ {
			nums[i], nums[level] = nums[level], nums[i]
			dfs(level + 1)
			nums[i], nums[level] = nums[level], nums[i]
		}
	}
	dfs(0)
	return res
}
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
