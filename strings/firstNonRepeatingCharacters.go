package strings

/*
 * Observation - We need to return the index of first non-repeating character, if not found then return -1
 * Brute force Approach - In this approach,
 	- Need to traverse over string
 	- Initialize a variable called foundDuplicate with boolean value assigned
	- Traverse over string again
	- Check if both char1 == char2 && id1 != id2
	- foundDuplicate updates to true
	- Check if no duplicate then return index
 *
 * Efficient Approach -
 *
 * T - O(n) | S - O(1)
*/

/*
 - Params -> str string
 - Return type => integer
 - T - O(n^2) | S - O(1)
*/
func NaiveFirstNonRepeatingCharacter(str string) int {
	for idx := range str {
		var foundDuplicate = false
		for idx2, _ := range str {
			if str[idx] == str[idx2] && idx != idx2 {
				foundDuplicate = true
			}
		}

		if !foundDuplicate {
			return idx
		}
	}
	return -1
}
