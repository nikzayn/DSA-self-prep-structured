package arrays

/*
 * Description - To return the array with updated moved elements
 * @Input - arr -> array of integers, toMove -> integer
 * @Output - array of integers
 */
// T - O(n) | S - O(1)
func moveElementToEnd(arr []int, toMove int) []int {
	i, j := 0, len(arr)-1
	for i < j {
		for i < j && arr[j] == toMove {
			j--
		}
		if arr[i] == toMove {
			arr[i], arr[j] = arr[j], arr[i]
		}
		i++
	}
	return arr
}
