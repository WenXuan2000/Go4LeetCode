package main

import "fmt"

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			candies[i] = 1
		} else {
			if ratings[i] > ratings[i-1] && candies[i] <= candies[i-1] {
				candies[i] = candies[i-1] + 1
			} else if candies[i] == 0 {
				candies[i] = 1
			}
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		} else if candies[i] == 0 {
			candies[i] = 1
		}
	}
	sum := 0
	for i := 0; i < len(ratings); i++ {
		sum += candies[i]
	}
	return sum
}
func main() {
	rating := []int{1, 2, 87, 87, 87, 2, 1}
	fmt.Println(candy(rating))
}
