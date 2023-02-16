package arrays

/*
 * Description - Find second largest in an array
 * Approach - Intitialize first and second variable, then traverse over array
 * - update second to first and first to current index, and after that check if
 * - second is less than current index and first is not equal to current index
 * - then update second to currentIndex and return the second at last
 * TC  O(n) | SC - O(1)
 */

func SecondLargest(arr []int) int {
	first := arr[0]
	second := -1

	for i := 0; i < len(arr); i++ {
		if first < arr[i] {
			second = first
			first = arr[i]
		}
		if second < arr[i] && first != arr[i] {
			second = arr[i]
		}
	}
	return second
}
