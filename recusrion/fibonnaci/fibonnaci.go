package fibonnaci


func GetNthFib(n int) int {
	if(n == 1) {
		return 0
	}
	if(n == 2) {
		return 1
	}
	return GetNthFib(n - 2) + GetNthFib(n - 1)
}

func GetNthFibHashMap(n int) int {
	return helper(n, memoize[int]int{
		1: 0,
		2: 1,
	})
}

func helper(n int, memoize map[int]int) int {
	if val, found := memoize[n]; found {
		return val
	}
	memoize[n] = helper(n - 1, memoize) + helper(n - 2, memoize)
	return memoize[n]
}