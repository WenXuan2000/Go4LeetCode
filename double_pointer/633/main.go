package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	a := 0
	b := int(math.Sqrt(float64(c)))
	for b >= a {
		if a*a+b*b > c {
			b -= 1
		} else if a*a+b*b < c {
			a += 1
		} else if a*a+b*b == c {
			return true
		}
	}
	return false
}
func main() {
	c := 2
	fmt.Println(judgeSquareSum(c))
}
