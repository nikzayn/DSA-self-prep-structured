package arrays

/*
 * Observation - To find the first duplicate value from the array
 * Approach -
   * Traverse over the array
   *
*/

func firstDuplicateValue(array []int) int {
	minimumIndex := len(array)
	for i := 0; i < len(array); i++ {
		firstVal := array[i]
		for j := i + 1; j < len(array); j++ {
			valueCompare := array[j]
			if firstVal == valueCompare {
				minimumIndex = min(minimumIndex, i)
			}
		}
	}
	if minimumIndex == len(array) {
		return -1
	}
	return array[minimumIndex]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
