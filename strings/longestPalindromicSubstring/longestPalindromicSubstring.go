package strings

func longestPalindromicSubstring(str string) string {
	longest := ""
	for i := range str {
		for j := i; j < len(str); j++ {
			substring := str[i : j+1]
			if len(substring) > len(longest) && IsPalindrome(substring) {
				longest = substring
			}
		}
	}
	return longest
}

func IsPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - i - 1
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
