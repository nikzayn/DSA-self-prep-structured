package arrays

/**
 * Descripion - To find the missing number from thr range of numbers
 * @Input() -> array of integers
 * @Output() -> integer
 * TC - O(N) | SC - O(N)
 */
func findMissingNum(arr []int) int {
	n := len(arr) + 1
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, num := range arr {
		actualSum += num
	}

	return expectedSum - actualSum
}
