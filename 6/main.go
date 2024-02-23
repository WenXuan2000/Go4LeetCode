package main

import "fmt"

func convert(s string, numRows int) string {
	matrix := make([][]byte, numRows)
	if numRows == 1 {
		return s
	}
	if len(s) == 0 {
		return s
	}
	tocol := (len(s) / (2*numRows - 2)) * (numRows - 1)
	if len(s)%(2*numRows-2) != 0 {
		if tocol == 0 {
			tocol += 1 + (len(s)%(2*numRows-2))%numRows
		} else if (len(s)%(2*numRows-2))/numRows >= 1 {
			tocol += 1 + (len(s)%(2*numRows-2))%numRows
		} else {
			tocol += 1
		}
	}
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, tocol)
	}
	col := 0
	row := 0
	for i := 0; i < len(s); i++ {
		matrix[row][col] = s[i]
		if (i+1)%(2*numRows-2) >= numRows || (i+1)%(2*numRows-2) == 0 || row == numRows-1 {
			col += 1
			row -= 1
		} else {
			row += 1
		}
	}
	ans := []uint8{}
	for i := 0; i < numRows; i++ {
		for j := 0; j < tocol; j++ {
			if matrix[i][j] != 0 {
				ans = append(ans, matrix[i][j])
			}
		}
	}
	return string(ans)
}
func main() {
	s := "PAYPALISHIRING"
	numRows := 5
	fmt.Println(convert(s, numRows))
}
