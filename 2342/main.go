package main

import (
	"fmt"
)

func maximumSum(nums []int) int {
	index := make(map[int][]int)
	ans := -1
	for i := 0; i < len(nums); i++ {
		ns := 0
		for ntemp := nums[i]; ntemp > 0; ntemp /= 10 {
			ns += ntemp % 10
		}
		if v, ok := index[ns]; ok {
			if nums[i] > v[0] && v[1] != -1 {
				biger := v[1] - v[0]
				ans = max(ans, biger+nums[i])
				if biger > nums[i] {
					index[ns] = []int{nums[i], biger + nums[i]}
				} else {
					index[ns] = []int{biger, biger + nums[i]}
				}

			} else if v[1] == -1 {
				ans = max(ans, v[0]+nums[i])
				if v[0] > nums[i] {
					index[ns] = []int{nums[i], v[0] + nums[i]}
				} else {
					index[ns] = []int{v[0], v[0] + nums[i]}
				}
			}
		} else {
			index[ns] = []int{nums[i], -1}
		}
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}
func main() {
	nums := []int{314593573, 110074004, 699559522, 673146088, 942819791, 869240392, 882704685, 539112585, 403921998, 255500178, 844178479, 767111119, 892897819, 885553426, 992987422, 914189130, 977810434, 152598980, 453977488, 389624268, 428531796, 839330802, 310252480, 378105000, 847373518, 934908066, 401794474, 663507194, 319531245, 614248496, 887058076, 278608939, 932087968, 282329958, 863096195, 98031682, 778619077, 471900584, 647816311, 469918315, 168055925, 550222361, 650029951, 280019987, 600359910, 803792276}
	fmt.Println(maximumSum(nums))
}

// 835
