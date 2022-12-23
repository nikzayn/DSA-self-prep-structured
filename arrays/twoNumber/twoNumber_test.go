package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	array    []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []addTest{
		{
			array:    []int{3, 5, -4, 8, 11, 1, -1, 6},
			target:   10,
			expected: []int{11, -1},
		},
		{
			array:    []int{4, 6},
			target:   10,
			expected: []int{4, 6},
		},
		{
			array:    []int{4, 6, 1, -3},
			target:   3,
			expected: []int{6, -3},
		},
		{
			array:    []int{4, 6, 1},
			target:   5,
			expected: []int{4, 1},
		},
	}
	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - TwoSum should return expected output", ix), func(t *testing.T) {
			output := TwoNumberNaive(tc.array, tc.target)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
