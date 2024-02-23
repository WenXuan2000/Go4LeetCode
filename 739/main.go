package main

import "fmt"

//func dailyTemperatures(temperatures []int) []int {
//	ans := make([]int, len(temperatures))
//	stack := []int{}
//	for i := 0; i < len(temperatures); i++ {
//		for len(stack) > 0 {
//			top := stack[len(stack)-1]
//			if temperatures[i] <= temperatures[top] {
//				break
//			}
//			stack = stack[:len(stack)-1]
//			ans[top] = i - top
//		}
//		stack = append(stack, i)
//	}
//	return ans
//}
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := []int{}
	for i, x := range temperatures {
		for len(st) != 0 && x > temperatures[st[len(st)-1]] {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
