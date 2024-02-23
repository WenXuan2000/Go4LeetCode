package main

import "fmt"

func intToRoman(num int) string {
	index := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	val := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := []byte{}
	for i := 0; i < len(index); i++ {
		if time := num / index[i]; time >= 1 {
			for j := 0; j < time; j++ {
				res = append(res, val[i]...)
			}
			num -= index[i] * time
		}
		if num == 0 {
			break
		}
	}
	return string(res)
}
func main() {
	fmt.Println(intToRoman(1994))
}
