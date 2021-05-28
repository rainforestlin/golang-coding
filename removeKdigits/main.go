package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	var stack []byte
	if len(num) < k {
		return " "
	}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

func main() {
	fmt.Println(removeKdigits("432502", 8))
}
