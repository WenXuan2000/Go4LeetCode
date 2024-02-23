package main

import "fmt"

func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	right[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		if left[i-1] > height[i] {
			left[i] = left[i-1]
		} else {
			left[i] = height[i]
		}
	}
	for i := len(height) - 2; i >= 0; i-- {
		if right[i+1] > height[i] {
			right[i] = right[i+1]
		} else {
			right[i] = height[i]
		}
	}
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		sum += Min(left[i], right[i]) - height[i]
	}
	return sum
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}
