package strings

import "strings"

func ReverseString(str string) string {
	//split the string
	words := strings.Fields(str)

	//Reverse the order of the words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	//Join the words
	return strings.Join(words, " ")
}
