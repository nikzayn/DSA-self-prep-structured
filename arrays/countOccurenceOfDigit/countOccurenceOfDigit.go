package arrays

func countDigit2(n int) int {
	if n < 2 {
		return 0
	}
	count := 0
	for i := 2; i <= n; i++ {
		count += count2IntDigit(i)
	}
	return count
}

func count2IntDigit(num int) int {
	count := 0
	for num > 0 {
		if num%10 == 2 {
			count++
		}
		num /= 10
	}
	return count
}
