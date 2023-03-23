package main

import "fmt"

func beans(height []int) (ans int) {
	/*
		单调栈做法
	*/
	// 初始化 栈，栈里面保存的是index
	var stack []int
	// 遍历height
	for i, h := range height {
		// 栈的大小要大于0，现有的高度要高于栈顶元素
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			// 获得栈顶元素index
			top := stack[len(stack)-1]
			// 出栈操作
			stack = stack[:len(stack)-1]
			// 如果栈空了，就结束循环；
			if len(stack) == 0 {
				break
			}
			// 获得`|_|` `左边的`|`
			left := stack[len(stack)-1]
			// 获得矩形的宽度
			w := i - 1 - left
			// 获得矩形的高度 取左边和右边小的柱子高度
			h := min(height[left], h) - height[top]
			ans += w * h
		}
		stack = append(stack, i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{5, 0, 2, 1, 4, 0, 1, 0, 3}
	fmt.Printf("赞青豆，赞青豆，看看你有多少青豆，原来你有%d个青豆", beans(height))
}

/*
时间复杂度：O(n)，其中 n 是切片 height 的长度。从 0 到 n-1 的每个下标最多只会入栈和出栈各一次。

空间复杂度：O(n)，其中 n 是切片 height 的长度。空间复杂度主要取决于栈空间，栈的大小不会超过 n。

*/
