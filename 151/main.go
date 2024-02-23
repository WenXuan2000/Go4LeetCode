package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	b := []string{}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != "" {
			b = append(b, a[i])
		}
	}
	return strings.Join(b, " ")
}
func main() {
	fmt.Println(reverseWords("  hello world  "))
}
