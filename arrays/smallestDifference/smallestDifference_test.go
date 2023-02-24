package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	arg2     []int
	expected []int
}

func TestSortedSquares(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{-1, 5, 10, 20, 28, 3},
			arg2:     []int{26, 134, 135, 15, 17},
			expected: []int{28, 26},
		},
		{
			arg1:     []int{-1, 5, 10, 20, 3},
			arg2:     []int{26, 134, 135, 15, 17},
			expected: []int{20, 17},
		},
		{
			arg1:     []int{10, 0, 20, 25},
			arg2:     []int{1005, 1006, 1014, 1032, 1031},
			expected: []int{25, 1005},
		},
		{
			arg1:     []int{10, 0, 20, 25, 2200},
			arg2:     []int{1005, 1006, 1014, 1032, 1031},
			expected: []int{25, 1005},
		},
		{
			arg1:     []int{240, 124, 86, 111, 2, 84, 954, 27, 89},
			arg2:     []int{1, 3, 954, 19, 8},
			expected: []int{954, 954},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - sorted squares should return expected output", idx), func(t *testing.T) {
			output := SmallestDifference(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
