package arrays

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	totalWaitingTime := 0
	for idx, duration := range queries {
		queriesLeft := len(queries) - (idx + 1)
		totalWaitingTime += duration * queriesLeft
	}
	return totalWaitingTime
}
