package eval_input

import (
	"fmt"
	"strings"
)

func Eval_input(input string) {

	// 使用 strings.Replace 函数将中括号替换为大括号
	output := strings.ReplaceAll(input, "[", "{")
	output = strings.ReplaceAll(output, "]", "}")

	fmt.Println(output)
}
