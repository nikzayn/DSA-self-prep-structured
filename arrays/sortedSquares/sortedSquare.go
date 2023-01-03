/*
 ** Observation - We have a sorted array and we need to return new array with same length but each index should be a squared value.
 *
 ** Scenarios - Check if array is sorted or not
 * Check if values are negative or not
 * If value is negative we need to implement modulus for that.
 *
 ** Approach
 * Naive - Initialize an empty slice and then iterate over an array and update the squared value to the particular index with val * val
 * after that sort the result because we need to return in ascending/sorted format array.
 */
package arrays

import "sort"

// T - O(nlogn) | S - O(n)
/*
 * Description - To return the squared value of each index in an array in ascending/sorted format.
 * Input - arr -> array of integers
 * Output - array of integers
 */
func sortedSquaresNaive(arr []int) []int {
	sortedArr := make([]int, len(arr))

	for idx, val := range arr {
		sortedArr[idx] = val * val
	}

	sort.Ints(sortedArr)
	return sortedArr
}

// T - O(n) | S - O(n)
/*
 * Description - To return the squared value of each index in an array in ascending/sorted format.
 * Approach - Two pointer approach with start and end index
 * Input - arr -> array of integers
 * Output - array of integers
 */
func sortedSquaresEff(arr []int) []int {
	sortedArr := make([]int, len(arr))
	start := 0
	end := len(arr) - 1

	for i := len(arr) - 1; i >= 0; i-- {
		smallVal := arr[start]
		largestVal := arr[end]

		if abs(smallVal) > abs(largestVal) {
			sortedArr[i] = smallVal * smallVal
			start += 1
		} else {
			sortedArr[i] = largestVal * largestVal
			end -= 1
		}
	}
	return sortedArr
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
