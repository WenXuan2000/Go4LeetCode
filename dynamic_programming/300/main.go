package main

import (
	"fmt"
)

// DP方法
//func lengthOfLIS(nums []int) int {
//	dp := make([]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		dp[i] = 1
//		for j := 0; j < i; j++ {
//			if nums[j] < nums[i] {
//				dp[i] = Max(dp[i], dp[j]+1)
//			}
//		}
//	}
//	max := 0
//	for _, v := range dp {
//		if v > max {
//			max = v
//		}
//	}
//	return max
//}
//二分查找加速
//func lengthOfLIS(nums []int) int {
//	g := []int{}
//	for _, x := range nums {
//		j := sort.SearchInts(g, x)
//		if j == len(g) { // >=x 的 g[j] 不存在
//			g = append(g, x)
//		} else {
//			g[j] = x
//		}
//	}
//	return len(g)
//}
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		if len(dp) == 0 || dp[len(dp)-1] < num {
			dp = append(dp, num)
		} else {
			l, r := 0, len(dp)-1
			pos := r
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = num
		} //二分查找
	}
	return len(dp)
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}
func main() {
	nums := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(lengthOfLIS(nums))
}
