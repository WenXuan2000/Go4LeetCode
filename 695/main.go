package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	// 每相邻两位即为上下左右四个方向之一

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				max = Max(max, dfs(grid, i, j))
			}
		}
	}
	return max
}
func dfs(grid [][]int, r, c int) int {
	direction := []int{-1, 0, 1, 0, -1}
	if grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	x, y := 0, 0
	area := 1
	for i := 0; i < 4; i++ {
		x = r + direction[i]
		y = c + direction[i+1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			area += dfs(grid, x, y)
		}
	}
	return area
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))
}
