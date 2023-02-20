package strings

//Observation - To find out maximum consecutive repeating characters in a string. Ex - 'AAaaabbbbaa'

/*
  * @Description - Initialize a maxCount, count and prevChar variable where ->
     * maxCount - 0
	 * currentCount - 1
	 * prevChar - To first character of the string
  * Traverse over length of string
  * Check if current element is equal to prevChar, if it is then increment the count
  * Else, check if count is greater than maxCount, then update maxCount to count,
  * Else reset the count to 1 and update the prevChar to current element
  * When iteration over length of string got completed, then check if count is greater
  * then maxCount, update maxCount to count
  * return maxCount
  * @Input - str -> string
  * @Output - integer
  * TC - O(N) | SC - O(N)
*/
func maximumConsecutiveRepeating(str string) int {
	maxCount := 0
	count := 0
	prevChar := str[0]

	for i := 1; i < len(str); i++ {
		if str[i] == prevChar {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 1
			prevChar = str[i]
		}
	}

	if count > maxCount {
		maxCount = count
	}

	return maxCount
}
