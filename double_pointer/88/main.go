package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	for m >= 1 || n >= 1 {
		if m == 0 {
			for n >= 1 {
				nums1[m] = nums2[m]
				n--
				m++
			}
			break
		} else if n == 0 {
			break
		}
		if nums2[n-1] >= nums1[m-1] {
			nums1[pos] = nums2[n-1]
			n -= 1
			pos -= 1
		} else if nums2[n-1] < nums1[m-1] {
			nums1[pos] = nums1[m-1]
			m -= 1
			pos -= 1
		}
	}
	fmt.Println(nums1)
}

func main() {
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//m := 3
	//nums2 := []int{2, 5, 6}
	//n := 3
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)
}
