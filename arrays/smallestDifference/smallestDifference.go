package arrays

import (
	"math"
	"sort"
)

/*
  * Observation - To find out the lowest absolute difference between two arrays elements
  * Approach -
    * Sort both arrays
	* Initialize index variables idxOne, idxTwo
	* Initialize smallest and current with MaxInt32
	* Initialize smallestPair slice
	* Check if ids is less than with both arrays
	* Check if first index element is less than second index element, update current = first - second and increment the idxOne
	* Check if first index element is greater than second index element, update current = second - first and increment the idxTwo
	* else return array with first and second
	* check if smallest is greater than current, if it is then update smallest with current and update the smallestPair with first and second
	* return smallestPair
	* TC - O(nlogn) + O(mlogn) | SC - O(1)
*/
func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	idxOne, idxTwo := 0, 0
	smallest, current := math.MaxInt32, math.MaxInt32
	smallestPair := []int{}

	for idxOne < len(array1) && idxTwo < len(array2) {
		first, second := array1[idxOne], array2[idxTwo]
		if first < second {
			current = second - first
			idxOne += 1
		} else if first > second {
			current = first - second
			idxTwo += 1
		} else {
			return []int{first, second}
		}

		if smallest > current {
			smallest = current
			smallestPair = []int{first, second}
		}
	}

	return smallestPair
}
