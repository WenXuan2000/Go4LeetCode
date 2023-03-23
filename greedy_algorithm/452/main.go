package main

import "sort"

//452. 用最少数量的箭引爆气球
//https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/

// 考虑右端点的情况
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}

//考虑左端点的情况
//func findMinArrowShots(points [][]int) int {
//	if len(points) == 0 {
//		return 0
//	}
//	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
//	ans := 1
//	for i:=1;i<len(points);i++{
//		if points[i][0] > points[i-1][1]{
//			ans++
//		}else{
//		// att：这里应该是取右端最短的
//			points[i][1] = min(points[i][1], points[i-1][1])
//		}
//	}
//	return ans
//}
//
//func min(x, y int) int{
//	if x > y{
//		return y
//	}
//	return x
//}

func main() {
	//points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	println(findMinArrowShots(points))
}
