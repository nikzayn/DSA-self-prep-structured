package arrays

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	triplets := [][]int{}

	for i := 0; i < len(array)-2; i++ {
		left := i + 1
		right := len(array) - 1

		if left < right {
			currentSum := array[i] + array[left] + array[right]

			if currentSum == target {
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if currentSum < target {
				left += 1
			} else if currentSum > target {
				right -= 1
			}
		}
	}
	return triplets
}
