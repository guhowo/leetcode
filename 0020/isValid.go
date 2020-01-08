package _020

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		size := len(stack)
		if size > 1 {
			top := stack[size-2]
			if top == ']' || top == ')' || top == '}' {
				return false
			}
			if s[i] == ']' && stack[size-2] == '[' {
				stack = stack[:size-2]
			}
			if s[i] == ')' && stack[size-2] == '(' {
				stack = stack[:size-2]
			}
			if s[i] == '}' && stack[size-2] == '{' {
				stack = stack[:size-2]
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
