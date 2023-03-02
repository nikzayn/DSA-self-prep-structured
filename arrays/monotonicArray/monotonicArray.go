package arrays

/*
 * description - Find an array which is monotonic in nature means an array who is non-decreasing and non-increasing in nature
 * @Input() - array -> Integer
 * @Output() - Boolean
 * TC - O(N) | SC - O(1)
 */
func isMontonicArray(array []int) bool {
	isNonDecreasing := true
	isNonIncreasing := true

	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNonDecreasing = false
		}
		if array[i] > array[i-1] {
			isNonIncreasing = false
		}
	}
	return isNonDecreasing || isNonIncreasing
}
