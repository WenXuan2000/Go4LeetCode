package main

func validPalindrome(s string) bool {
	n := len(s)
	if n < 2 {
		return true
	}

	// isPalindrome 验证s[begin, end]是否是回文字符串
	var isPalindrome func(string, int, int) bool
	isPalindrome = func(s string, begin, end int) bool {
		for begin < end {
			if s[begin] != s[end] {
				return false
			}
			begin++
			end--
		}
		return true
	}

	left, right := 0, n-1
	for left < right {
		if s[left] == s[right] { // 1. 左右指针的元素相同时，继续两端夹逼
			left++
			right--
		} else { // 2. 左右指针元素不同时，判断 s(left, right] 或 s[left, right) 是否是回文串
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
	}
	return true
}

func main() {
	s := "cuppucu"
	println(validPalindrome(s))
}
