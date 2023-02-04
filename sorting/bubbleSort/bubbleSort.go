package sorting

/*
 ** Observation - Need to bubble sort an array
 ** Approach -
 * Check if array is sorted or not with isSorted
 * Traverse over array
 * Update isSorted
 * Traverse over the array again while deducting the counter from it
 * Swap the values
 */

// T - O() | S - O(1)
/*
 * Description - To return the sorted array
 * Input - array -> array of integers
 * Output - array -> array of integers
 */

func bubbleSort(array []int) []int {
	isSorted := false
	counter := 0

	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				swap(array, i, i+1)
				isSorted = false
			}
		}
		counter++
	}
	return array
}

func swap(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}
