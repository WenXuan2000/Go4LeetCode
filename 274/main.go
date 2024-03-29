package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) (h int) {
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations))
}
