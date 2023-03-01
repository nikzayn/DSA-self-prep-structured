package main

import "fmt"

//Two pointer approach
func moveElementsToEnd(array []int, toMove int) []int {
	end := len(array) - 1
	for i := 0; i <= end; i++ {
		if array[i] == toMove {
			//swap
			array[i], array[end] = array[end], array[i]
			end--
			i--
		}
	}
	return array
}

func main() {
	arr := []int{1, 2, 1, 2, 2, 2, 3, 4, 2}
	fmt.Println(moveElementsToEnd(arr, 2))
}
