package validParentheses

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, v := range s {
		switch v {
		case '{', '[', '(':
			stack = append(stack, byte(v))
			continue
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack) - 1]
			if v == '}' && top != '{' ||
				v == ']' && top != '[' ||
				v == ')' && top != '(' {
				return false
			}

			stack = stack[:len(stack) - 1]
		}
	}

	return len(stack) == 0
}
