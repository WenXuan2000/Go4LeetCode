package main

import "fmt"

//https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	begin := 0
	end := len(numbers) - 1
	for i := 0; i < len(numbers); i++ {
		res := numbers[begin] + numbers[end]
		if res == target {
			return []int{begin + 1, end + 1}
		} else if res > target {
			end -= 1
		} else if res < target {
			begin += 1
		}
	}
	return []int{begin + 1, end + 1}
}
func main() {
	//number := []int{2, 7, 11, 15}
	//target := 9
	//number := []int{2, 3, 4}
	//target := 6
	number := []int{-1, 0}
	target := -1
	fmt.Println(twoSum(number, target))
}
