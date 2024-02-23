package main

import "fmt"

//func longestCommonPrefix(strs []string) string {
//	ans := 0
//	next := strs[0]
//	for i := 1; i < len(strs); i++ {
//		ans = judge(next, strs[i])
//		next = next[:ans]
//		if ans == 0 {
//			break
//		}
//	}
//	return next
//}
//func judge(a string, b string) int {
//	index := 0
//	for i := 0; i < len(a); i++ {
//		if i < len(b) {
//			if a[i] == b[i] {
//				index += 1
//				continue
//			}
//		}
//		break
//	}
//	return index
//}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := range strs[0] {
		for _, s := range strs[1:] {
			if i >= len(s) || s[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
