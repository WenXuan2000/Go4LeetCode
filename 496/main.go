package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	mp := map[int]int{}
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}

	for i, x := range nums2 {
		for len(st) != 0 && x > nums2[st[len(st)-1]] {
			mp[nums2[st[len(st)-1]]] = x
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	for i, x := range nums1 {
		ans[i] = mp[x]
		if ans[i] == 0 {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
