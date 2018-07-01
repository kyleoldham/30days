package main

func isValid(s string) bool {
	// Declare mapping between (), {}, and []
	// Use a stack as the method of checking

	stk := new(stack)

	for _, b := range s {
		switch b {
		case '(', '[', '{':
			stk.Push(b)
		case ')', ']', '}':
			if res, ok := stk.Pop(); !ok || res != twinning[b] {
				return false
			}
		}
	}

	// If stack not empty
	if len(*stk) > 0 {
		return false
	}

	// Everything Matched
	return true
}

// These need to be opposite than expected, cuz it's looking for closing
var twinning = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

type stack []rune

func (stk *stack) Push(b rune) {
	*stk = append(*stk, b)
}

func (stk *stack) Pop() (rune, bool) {
	if len(*stk) > 0 {
		res := (*stk)[len(*stk)-1]
		*stk = (*stk)[:len(*stk)-1]
		return res, true
	}
	return 0, false
}
