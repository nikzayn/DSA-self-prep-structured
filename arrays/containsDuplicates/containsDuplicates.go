package arrays

import "fmt"

// TC - O(n^2) | SC - O(1)
func containsDuplicates(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsDuplicatesEff(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, found := set[num]; found {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicatesEff(nums))
}
