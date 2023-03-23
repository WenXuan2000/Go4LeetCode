package main

import "math"

func minWindow(s string, t string) string {
	//ori和cnt是两张哈希表
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				//记录最短长度
				len = r - l + 1
				//记录最短长度时的左边界和右边界
				ansL, ansR = l, l+len
			}
			//
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	println(minWindow(s, t))
}
