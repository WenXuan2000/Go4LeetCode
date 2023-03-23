package main

func findLongestWord(s string, dictionary []string) string {
	maxlen := 0
	var ans string
	check := func(tar string, src string) bool {
		ind1, ind2 := 0, 0
		for ind1 < len(tar) && ind2 < len(src) {
			if tar[ind1] == src[ind2] {
				ind1++
				ind2++
			} else {
				ind2++
			}
		}
		if ind1 != len(tar) {
			return false
		} else {
			return true
		}
	}
	for _, src := range dictionary {
		if check(src, s) && len(src) > maxlen {
			maxlen = len(src)
			ans = src
		} else if check(src, s) && len(src) == maxlen && (src < ans || ans == "") {
			maxlen = len(src)
			ans = src
		}
	}
	return ans
}

func main() {
	s := "abce"
	dict := []string{"abe", "abc"}
	println(findLongestWord(s, dict))
}
