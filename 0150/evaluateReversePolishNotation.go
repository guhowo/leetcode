package _150

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	cur := 0
	for _, v := range tokens {
		//如果是运算符，那么取出stack中两个top，进行运算
		if v == "+" || v == "-" || v == "*" || v == "/" {
			before, after := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if v == "+" {
				cur = before + after
			} else if v == "-" {
				cur = before - after
			} else if v == "*" {
				cur = before * after
			} else if v == "/" {
				cur = before / after
			}
			stack = append(stack, cur)
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}
