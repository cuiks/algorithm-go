package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0
			}
			v1 := stack[len(stack)-1]
			v2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if v == "+" {
				stack = append(stack, v1+v2)
			}
			if v == "-" {
				stack = append(stack, v2-v1)
			}
			if v == "*" {
				stack = append(stack, v2*v1)
			}
			if v == "/" {
				stack = append(stack, v2/v1)
			}
		default:
			vInt, _ := strconv.Atoi(v)
			stack = append(stack, vInt)
		}
	}
	return stack[0]
}

func main() {

}
