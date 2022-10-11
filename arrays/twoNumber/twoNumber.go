package arrays

func TwoNumber(arr []int, target int) []int {
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
