package FastAndSlowPointers

/*
 * Observation - To find out the happy number if it is returning 1, else it's not a happy number
 * Approach -
	* Initialize a slow with input number and fast with squared of sum of digits
*/

func HappyNumber(n int) bool {
	slow, fast := n, n
	for {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))

		if slow == 1 || fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
