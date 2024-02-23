package main

import "fmt"

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m := len(heights)
	n := len(heights[0])
	p := make([][]bool, m)
	a := make([][]bool, m)
	for i := range p {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}
	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		for _, d := range dirs {
			if nx, ny := x+d.x, y+d.y; 0 <= nx && nx < m && 0 <= ny && ny < n && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, p)
	}
	for j := 1; j < n; j++ {
		dfs(0, j, p)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, a)
	}
	for j := 0; j < n-1; j++ {
		dfs(m-1, j, a)
	}
	for i, row := range p {
		for j, ok := range row {
			if ok && a[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
func main() {
	heights := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	fmt.Println(pacificAtlantic(heights))
}
