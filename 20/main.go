package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	parentheses := []byte{'(', ')', '{', '}', '[', ']'}
	for j := 0; j < len(s); j++ {
		for i := 0; i < len(parentheses); i++ {
			if s[j] == parentheses[i] {
				if i%2 == 0 {
					stack = append(stack, s[j])
					if len(stack) >= len(s)-j {
						return false
					}
				} else {
					if len(stack) == 0 {
						return false
					}
					top := stack[len(stack)-1]
					if top != parentheses[i-1] {
						return false
					} else {
						stack = stack[:len(stack)-1]
						break
					}
				}
			}
		}
	}
	return true
}
func main() {
	s := "["
	fmt.Println(isValid(s))
}
