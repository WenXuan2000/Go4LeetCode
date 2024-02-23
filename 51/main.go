package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	//r := make([]bool, n)

	res := [][]string{}
	ans := []string{}
	if n == 1 {
		ans = append(ans, "Q")
		copyAns := make([]string, n)
		copy(copyAns, ans)
		res = append(res, copyAns)
		return res
	}
	c := make([]bool, n)
	ld := make([]bool, 2*n-1)
	rd := make([]bool, 2*n-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			copyAns := make([]string, n)
			copy(copyAns, ans)
			res = append(res, copyAns)
			return
		}
		for col := 0; col < n; col++ {
			if !c[col] && !ld[n-r+col-1] && !rd[r+col] {
				ans = append(ans, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
				c[col] = true
				ld[n-r+col-1] = true
				rd[r+col] = true

				dfs(r + 1)

				c[col] = false
				ld[n-r+col-1] = false
				rd[r+col] = false
				ans = ans[:len(ans)-1]
			}
		}
	}
	dfs(0)
	return res
}
func main() {
	n := 1
	fmt.Println(solveNQueens(n))
}
