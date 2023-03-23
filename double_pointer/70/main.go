package main

import "fmt"

func climbStairs(n int) int {
	//dp := []int{}
	pre1 := 1
	pre2 := 2
	var temp int
	if n <= 2 {
		return n
	}
	for i := 2; i < n; i++ {
		temp = pre1 + pre2
		pre1 = pre2
		pre2 = temp
	}
	return temp
}
func main() {
	n := 4
	fmt.Println(climbStairs(n))
}
