package main

import (
	"fmt"
	"strings"
)

//func lengthOfLastWord(s string) (ans int) {
//	index := len(s) - 1
//	for s[index] == ' ' {
//		index--
//	}
//	for index >= 0 && s[index] != ' ' {
//		ans++
//		index--
//	}
//	return
//}
func lengthOfLastWord(s string) int {
	b := strings.TrimSpace(s)
	a := strings.Split(b, " ")
	return len(a[len(a)-1])
}
func main() {
	fmt.Println(lengthOfLastWord("Today is a nice day"))
}
