package main

import (
	"fmt"
	"math"
)

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i == m-1 {
			return grid[i][j]
		}
		p := &memo[i][j]
		if *p != 0 {
			return *p
		}
		res := math.MaxInt
		for k, c := range moveCost[grid[i][j]] {
			res = min(res, dfs(i+1, k)+c)
		}
		*p = res + grid[i][j]
		return res + grid[i][j]
	}
	ans := math.MaxInt
	for j := 0; j < n; j++ {
		ans = min(ans, dfs(0, j))
	}
	return ans
}
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}

}
func main() {
	grid := [][]int{{5, 3}, {4, 0}, {2, 1}}
	moveCost := [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}
	fmt.Println(minPathCost(grid, moveCost))
}
