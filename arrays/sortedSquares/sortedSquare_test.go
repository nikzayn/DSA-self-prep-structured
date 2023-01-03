package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	expected []int
}

func TestSortedSquares(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{1, 4, 9, 16, 25, 36, 49, 64},
		},
		{
			arg1:     []int{1},
			expected: []int{1},
		},
		{
			arg1:     []int{0},
			expected: []int{0},
		},
		{
			arg1:     []int{-2, -1},
			expected: []int{1, 4},
		},
		{
			arg1:     []int{-5, -4, -3, -2, -1},
			expected: []int{1, 4, 9, 16, 25},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - sorted squares should return expected output", idx), func(t *testing.T) {
			output := sortedSquaresEff(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
