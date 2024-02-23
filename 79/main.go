package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	visit := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visit[i] = make([]bool, len(board[0]))
	}
	index := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	var dfs func(int, int, int) bool
	dfs = func(r int, c int, k int) bool {
		//if res[len(res)-1] == word[len(res)-1] && (r != 0 && c != 0) {
		//	return
		//}
		//if board[r][c]==word[len(res)]{
		//	return
		//}
		//else if board[r+1][c]==word[len(res)]{
		//
		//}
		if board[r][c] != word[k] || visit[r][c] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visit[r][c] = true
		defer func() { visit[r][c] = false }()
		for _, v := range index {
			if r+v[0] >= 0 && r+v[0] < len(board) && c+v[1] >= 0 && c+v[1] < len(board[0]) {
				if dfs(r+v[0], c+v[1], k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
