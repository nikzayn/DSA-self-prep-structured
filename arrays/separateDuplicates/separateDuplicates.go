package arrays

// T - O(N) | S - O(N)
func SeparateDuplicatesNaive(arr []int) int {
	check := make(map[int]int)
	res := make([]int, 0)

	for _, val := range arr {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}

	return len(res)
}
