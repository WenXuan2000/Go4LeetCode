package main

import "fmt"

func vowelStrings(words []string, left int, right int) int {
	//index := []string{"a", "e", "i", "o", "u"}
	//var set map[string]int
	//set = make(map[string]int)
	set := map[byte]int{
		'a': 1,
		'e': 2,
		'i': 3,
		'o': 4,
		'u': 5,
	}
	//wordset := map[string]int{}
	var isvowe func(string) bool
	isvowe = func(words string) bool {
		_, bok := set[words[0]]
		if !bok {
			return false
		}
		_, eok := set[words[len(words)-1]]
		if !eok {
			return false
		}
		//_, haveok := wordset[words]
		//if !haveok {
		//	wordset[words] = 1
		//	return true
		//}
		return true
	}
	ans := 0
	for i := 0; i < len(words); i++ {
		if i < left || i > right {
			continue
		}
		if isvowe(words[i]) {
			ans += 1
		}
	}
	return ans
}
func main() {
	words := []string{"ou", "e", "e"}
	left := 0
	right := 2
	fmt.Println(vowelStrings(words, left, right))
}
