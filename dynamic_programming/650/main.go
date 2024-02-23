package main

import (
	"fmt"
)

func minSteps(n int) int {
	f := make([]int, n+1)
	for i := 2; i <= n; i++ {
		f[i] = i
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				f[i] = Min(f[i], f[j]+i/j)
				f[i] = Min(f[i], f[i/j]+j)
			}
		}
	}
	return f[n]
}

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func main() {
	n := 3
	fmt.Println(minSteps(n))
}
