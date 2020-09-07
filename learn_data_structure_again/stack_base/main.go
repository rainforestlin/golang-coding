package main

import (
	"fmt"
)

var stack []byte

func isLeft(c byte) bool {
	if c == byte('{') || c == byte('[') || c == byte('(') {
		return true
	} else {
		return false
	}
}

func isPair(curr byte, pair byte) bool {
	if (curr == byte('}') && pair == byte('{')) || (curr == byte(']') && pair == byte('[')) || (curr == byte(')') && pair == byte('(')) {
		return true
	} else {
		return false
	}
}

func isLegal(s []byte) bool {
	for i := 0; i < len(s); i++ {
		curr := s[i]
		if isLeft(curr) {
			stack = append(stack, curr)
		} else {
			if len(stack) <= 0 {
				return false
			}
			pair := stack[len(stack)-1]
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			} else {
				stack = []byte{}
			}
			if !isPair(curr, pair) {
				return false
			}
		}
	}
	if len(stack) <= 0 {
		return true
	} else {
		return false
	}
}

func main() {
	input := "{[()()]}"
	var s = []byte(input)
	fmt.Println(isLegal(s))
}
