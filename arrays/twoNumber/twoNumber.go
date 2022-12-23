package arrays

// T - O(N^2) | S - O(1)
func TwoNumberNaive(arr []int, target int) []int {
	for i := 0; i < len(arr)-1; i++ {
		firstNum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			secondNum := arr[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

// T - O(N) | S - O(N)
func TwoNumberEff(arr []int, target int) []int {
	m := map[int]bool{}
	for _, num := range arr {
		potentialMatch := target - num
		if _, found := m[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		m[num] = true
	}
	return []int{}
}
