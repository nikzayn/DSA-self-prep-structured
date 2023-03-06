package strings

/*
 * Description - To find the longest substring without repeating characters
 * Approach -
	* Initialize a hashmap with byte and integer
	* Initialize a variable ans
	* Traverse over the string
	*
 * TC - O(N) | SC - O(N)
*/

func findLongestSubstring(str string) int {
	len := len(str)
	ans := 0
	m := make(map[byte]int)

	for i, j := 0, 0; j < len; j++ {
		//Check if index is found
		if index, ok := m[str[j]]; ok {
			i = max(index, i)
		}

		ans = max(ans, j-i+1)
		m[str[j]] = j + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
