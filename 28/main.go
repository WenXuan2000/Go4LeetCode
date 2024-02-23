package main

import "fmt"

//func strStr(haystack string, needle string) int {
//	for i := 0; i < len(haystack)-len(needle)+1; i++ {
//		if haystack[i] == needle[0] {
//			found := true // 设置标志位来表示是否找到匹配
//			for j := 1; j < len(needle); j++ {
//				if haystack[i+j] != needle[j] {
//					found = false
//					break
//				}
//			}
//			if found {
//				return i
//			}
//		}
//	}
//	return -1
//}
//KMP 算法
func strStr(haystack string, needle string) int {
	//next := make([]int, len(needle))
	//for j, i := 0, 1; i < len(needle); {
	//	if needle[j] == needle[i] {
	//		next[i] = j + 1
	//		i++
	//		j++
	//	} else {
	//		if j != 0 {
	//			j = next[j-1]
	//		} else {
	//			next[i] = 0
	//			i++
	//		}
	//	}
	//}
	next := make([]int, len(needle))
	for j, i := 0, 1; i < len(needle); {
		if needle[i] == needle[j] {
			next[i] = j + 1
			i++
			j++
		} else {
			if j != 0 {
				j = next[j-1]
			} else {
				next[i] = 0
				i++
			}
		}
	}
	for i, j := 0, 0; i < len(haystack); {
		if haystack[i] == needle[j] {
			j++
			i++
			if j == len(needle) {
				return i - j
			}
		} else {
			if j != 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}
func main() {
	haystack := "aabaaabaaac"
	needle := "aabaaac"
	fmt.Println(strStr(haystack, needle))
}
