package stacks

/*
  * Description - To remove adjacent duplicates from the string and return the string without duplicates
  ** Approach **
	* Initialize a stack to store characters and frequencies
	* Traverse over string, to check if current character is equal to the top of the stack elem
	* If yes, then pop the character from the stack
	* Else, append to the stack with the current character
	* return the stack
  * TC - O(N) | SC - O(N)
*/
func removeDuplicates(str string) string {
	stack := []rune{}

	//last := stack[len(stack)-1]
	for _, c := range str {
		if len(stack) > 0 && c == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}
