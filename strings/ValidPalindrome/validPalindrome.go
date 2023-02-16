package strings

// TC - O(N) | O(N)
func NaiveValidPalindrome(val string) bool {
	result := []byte{}
	for i := len(val) - 1; i >= 0; i-- {
		result = append(result, val[i])
	}
	return val == string(result)
}

func EffValidPalindrome(val string) bool {
	for i := 0; i < len(val); i++ {
		j := len(val) - 1 - i
		if val[i] != val[j] {
			return false
		}
	}
	return true
}
