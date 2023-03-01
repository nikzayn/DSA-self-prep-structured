package main

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	arg2     int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{2, 1, 2, 2, 2, 3, 4, 2},
			arg2:     2,
			expected: []int{4, 1, 3, 2, 2, 2, 2, 2},
		},
		{
			arg1:     []int{},
			arg2:     3,
			expected: []int{},
		},
		{
			arg1:     []int{1, 2, 4, 5, 6},
			arg2:     3,
			expected: []int{1, 2, 4, 5, 6},
		},
		{
			arg1:     []int{5, 5, 5, 5, 5, 5, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12},
			arg2:     5,
			expected: []int{12, 11, 10, 9, 8, 7, 1, 2, 3, 4, 6, 5, 5, 5, 5, 5, 5},
		},
	}
	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - TwoSum should return expected output", ix), func(t *testing.T) {
			output := moveElementsToEnd(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
