package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, v := range wordDict {
			length := len(v)
			if i < length {
				continue
			} else if s[i-length:i] == v {
				dp[i] = dp[i] || dp[i-length]
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "dogs"
	wordDict := []string{"dog", "s", "gs"}
	fmt.Println(wordBreak(s, wordDict))
}
