package validparentheses

func isValid(s string) bool {
	stack := []rune{}
	hashMap := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, c := range s {
		if idx, ok := hashMap[c]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == idx {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
