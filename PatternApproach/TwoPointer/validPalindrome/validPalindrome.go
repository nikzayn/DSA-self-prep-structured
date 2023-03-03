package TwoPointer

func ValidPalindrome(str string) bool {
	res := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}

	return str == string(res)
}

func EffValidPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
