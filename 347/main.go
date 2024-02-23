package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	frequence := make(map[int]int)
	for _, v := range nums {
		frequence[v]++
	}
	buckts := make([][]int, len(nums)+1)
	for num, freq := range frequence {
		buckts[freq] = append(buckts[freq], num)
	}
	res := make([]int, 0)
	for i := len(buckts) - 1; k > 0; i-- {
		for _, v := range buckts[i] {
			res = append(res, v)
			k--
			if k == 0 {
				break
			}
		}
	}
	return res
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
