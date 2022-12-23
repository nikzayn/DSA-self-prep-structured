package arrays

/*
 * Observation - We need to check that element in our main array should be existing in our sequence array.
 * If subsequence is not valid, then return false
 *
 * Approach - Initialize two variables to mark the main array and subsequence array counts.
 * Check with while loop that, each variable length should be less to both arrays
 * and inside it check if index of both arrays matches return true, else false

 * T - O(n) | S - O(1)
 */
func validSubsequence(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx += 1
		}
		arrIdx += 1
	}
	return seqIdx == len(sequence)
}
