package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	res := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := ")()())"
	fmt.Println(longestValidParentheses(s))
}
