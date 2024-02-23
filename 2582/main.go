package main

import "fmt"

func passThePillow(n int, time int) int {
	if n == 1 {
		return n
	}
	j := time / (n - 1)
	if j%2 == 0 {
		return time - j*(n-1) + 1
	} else {
		return n - (time - j*(n-1))
	}
}
func main() {
	fmt.Println(passThePillow(18, 38))
}
