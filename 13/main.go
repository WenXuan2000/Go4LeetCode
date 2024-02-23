package main

import "fmt"

func romanToInt(s string) int {
	ans := 0
	mymap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		if (s[i] == 'V' || s[i] == 'X') && i-1 >= 0 && s[i-1] == 'I' {
			ans += mymap[s[i]] - mymap[s[i-1]]*2
		} else if (s[i] == 'L' || s[i] == 'C') && i-1 >= 0 && s[i-1] == 'X' {
			ans += mymap[s[i]] - mymap[s[i-1]]*2
		} else if (s[i] == 'D' || s[i] == 'M') && i-1 >= 0 && s[i-1] == 'C' {
			ans += mymap[s[i]] - mymap[s[i-1]]*2
		} else {
			ans += mymap[s[i]]
		}
	}
	return ans
}
func main() {
	s := "IV"
	fmt.Println(romanToInt(s))
}
