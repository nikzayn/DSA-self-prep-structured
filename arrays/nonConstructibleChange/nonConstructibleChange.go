/*
 ** Observation - We have an array of values which denotes coins value and we need to check how many changes that we cannot create from the
 ** values available
 ** Scenarios - Check if array is sorted or not
 * Check if values are negative or not
 * If value is negative we need to implement modulus for that.
 *
 ** Approach -
 * sort the index values
 * Initialize a count variable
 * Iterate over an array and check if a coin value is greater than count + 1, if it is then return the count + 1
 * else update count value with the addition of coin value
 * return the value of count + 1 at last
 */

// T - O(nlogn) | S - O(1)
/*
 * Description - To return the count of non constructible change
 * Input - arr -> array of integers
 * Output - int
 */
package arrays

import "sort"

func nonConstructibleChange(arr []int) int {
	sort.Ints(arr)
	currentChangeValue := 0

	for _, coin := range arr {
		if coin > currentChangeValue+1 {
			return currentChangeValue + 1
		}
		currentChangeValue += coin
	}
	return currentChangeValue + 1
}
