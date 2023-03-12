package stacks

/*
 * Description - Given a string with matched and unmatched parentheses,
 remove the minimum number of parentheses so that the resulting string is a valid parenthesization.
 * Approach - Initialize an empty stack
	- Traverse over the string
	- Add the character if open brackets found,
	- Else, if Add the closing bracket if len of stack is zero, else pop the top from the stack
	- Initialize one empty variable with rune
	- Traverse over the stack, update the each result with * placeholder
	- Initialize an empty string variable, traverse over result, check if character != "*"
	- Then, concatenate the each character
	- return the string
 * @Input() - str -> string
 * @Output() -> string
 * TC - O(N) | O(N)
*/
func minimumRemoveParentheses(str string) string {
	stack := []int{}

	for i, c := range str {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	result := []rune(str)
	for _, i := range stack {
		result[i] = '*'
	}

	validString := ""
	for _, c := range result {
		if c != '*' {
			validString += string(c)
		}
	}
	return validString
}
