package main

import (
	"fmt"
	"sort"
)

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	sortedNums := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		sortedNums[i] = []int{nums1[i], nums2[i]}
	}
	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i][0] > sortedNums[j][0]
	})

	sortedQueries := make([][]int, len(queries))
	for i := 0; i < len(queries); i++ {
		sortedQueries[i] = []int{i, queries[i][0], queries[i][1]}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][1] > sortedQueries[j][1]
	})

	stack := [][]int{}
	answer := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		answer[i] = -1
	}
	j := 0
	for _, q := range sortedQueries {
		i, x, y := q[0], q[1], q[2]
		for j < len(sortedNums) && sortedNums[j][0] >= x {
			num1, num2 := sortedNums[j][0], sortedNums[j][1]
			for len(stack) > 0 && stack[len(stack)-1][1] <= num1+num2 {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1][0] < num2 {
				stack = append(stack, []int{num2, num1 + num2})
			}
			j++
		}
		k := sort.Search(len(stack), func(i int) bool {
			return stack[i][0] >= y
		})

		if k < len(stack) {
			answer[i] = stack[k][1]
		}
	}
	return answer
}

//type BinaryIndexedTree struct {
//	n   int
//	ans []int
//}
//
//func NewBIT(n int) BinaryIndexedTree {
//	ans := make([]int, n+1)
//	for i := range ans {
//		ans[i] = -1
//	}
//	return BinaryIndexedTree{n: n, ans: ans}
//}
//func (bit *BinaryIndexedTree) update(x, v int) {
//	for x <= bit.n {
//		bit.ans[x] = max(bit.ans[x], v)
//		x += x & -x
//	}
//}
//
//func (bit *BinaryIndexedTree) query(x int) int {
//	mx := -1
//	for x > 0 {
//		mx = max(mx, bit.ans[x])
//		x -= x & -x
//	}
//	return mx
//}
//func max(x, y int) int {
//	if x > y {
//		return x
//	} else {
//		return y
//	}
//}
//func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
//	n, m := len(nums1), len(queries)
//	nums := make([][2]int, n)
//	for i := range nums {
//		nums[i] = [2]int{nums1[i], nums2[i]}
//	}
//	sort.Slice(nums, func(i, j int) bool { return nums[j][0] < nums[i][0] })
//	sort.Ints(nums2)
//	idx := make([]int, m)
//	for i := range idx {
//		idx[i] = i
//	}
//	sort.Slice(idx, func(i, j int) bool { return queries[idx[j]][0] < queries[idx[i]][0] })
//	tree := NewBIT(n)
//	ans := make([]int, m)
//	j := 0
//	for _, i := range idx {
//		x, y := queries[i][0], queries[i][1]
//		for ; j < n && nums[j][0] >= x; j++ {
//			k := n - sort.SearchInts(nums2, nums[j][1])
//			tree.update(k, nums[j][0]+nums[j][1])
//		}
//		k := n - sort.SearchInts(nums2, y)
//		ans[i] = tree.query(k)
//	}
//	return ans
//}

func main() {
	nums1 := []int{4, 3, 1, 2}
	nums2 := []int{2, 4, 9, 5}
	queries := [][]int{{4, 1}, {1, 3}, {2, 5}}
	fmt.Println(maximumSumQueries(nums1, nums2, queries))
}
