package validParenthesis

/*
	* Description - To validate if all the parenthesis are open and closed properly
**Approach**
	* We can initialize stack data structure to track open parentheses
	* Traverse over the string
	* Check if opening parenthesis of any type are there, if there add them in stack
	* Else, if the character is closing parentheses, then
*/
func validParenthesis(str string) bool {
	stack := []rune{}

	for _, char := range str {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if char == '(' && last != ')' {
				return false
			} else if char == '{' && last != '}' {
				return false
			} else if char == '[' && last != ']' {
				return false
			}
		}
	}
	return len(stack) == 0
}
