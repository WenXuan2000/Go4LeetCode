package main

import "fmt"

func numDecodings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			//这里为啥不是f[i-1]+1,因为1 2 3 等价 1|2 3
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

func main() {
	s := "27"
	fmt.Println(numDecodings(s))
}
