package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	nums := []int{}
	var dfs func(int)
	dfs = func(index int) {
		if len(nums) == k {
			temp := []int{}
			temp = append(temp, nums...)
			res = append(res, temp)
			return
		}
		for i := index; i <= n; i++ {
			nums = append(nums, i)
			dfs(i + 1)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(1)
	return res
}
func main() {
	n, k := 4, 2
	fmt.Println(combine(n, k))
}
