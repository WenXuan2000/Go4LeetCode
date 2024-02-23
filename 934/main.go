package main

import (
	"fmt"
	"go4leetcode/eval_input"
)

func shortestBridge(grid [][]int) int {
	findone := false
	//visit := make([][]bool, len(grid))
	//for i := 0; i < len(grid); i++ {
	//	visit = append(visit, make([]bool, len(grid)))
	//}
	queue := [][]int{}
	var dfs func(int, int)
	dfs = func(r int, c int) {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == 2 {
			return
		}
		if grid[r][c] == 0 {
			queue = append(queue, []int{r, c})
		}
		if grid[r][c] == 1 {
			grid[r][c] = 2
			dfs(r+1, c)
			dfs(r, c+1)
			dfs(r-1, c)
			dfs(r, c-1)
		}
	}
	for i := 0; i < len(grid); i++ {
		if findone {
			break
		}
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				findone = true
				break
			}
		}
	}
	//	这里已经处理好了，有岛屿1 和2
	ans := 0
	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) > 0 {
		ans++
		n := len(queue)
		for i := n; i > 0; i-- {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				x, y := r+direction[k][0], c+direction[k][1]
				if x >= 0 && y >= 0 && x < len(grid) && y < len(grid) {
					if grid[x][y] == 2 {
						continue
					}
					if grid[x][y] == 1 {
						return ans
					}
					queue = append(queue, []int{x, y})
					grid[x][y] = 2
				}
			}
		}
	}
	return ans
}
func main() {
	eval_input.Eval_input("{{0,0,0,0,0,0,0},{1,0,0,1,0,0,0},{1,1,1,1,0,0,0},{0,0,0,1,0,0,0},{1,0,0,0,0,0,0},{1,1,0,0,0,0,0},{1,0,0,0,0,0,0}}")
	grid := [][]int{{0, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 1, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0}}
	fmt.Println(shortestBridge(grid))
}
