package amperity

func balanced(str string) bool {
	openClose := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	closeOpen := make(map[rune]rune)
	for k, v := range openClose {
		closeOpen[v] = k
	}

	var stack []rune
	for _, char := range str {
		if _, ok := openClose[char]; ok {
			stack = append(stack, char)
			continue
		}

		if openChar, ok := closeOpen[char]; ok {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != openChar {
				return false
			}

			stack = stack[:len(stack)-1]
		}

	}

	if len(stack) != 0 {
		return false
	}

	return true
}
