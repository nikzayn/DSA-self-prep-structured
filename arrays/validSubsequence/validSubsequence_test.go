package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	arg2     []int
	expected bool
}

func TestValidSubsequence(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{5, 1, 22, 25, 6, -1, 8, 10},
			arg2:     []int{1, 6, -1, 10},
			expected: true,
		},
		{
			arg1:     []int{5, 1, 22, 25, 6, -1, 8, 10},
			arg2:     []int{5, 1, 22, 23, 6, -1, 8, 10},
			expected: false,
		},
		{
			arg1:     []int{5, 1, 22, 25, 6, -1, 8, 10},
			arg2:     []int{22, 25, 6},
			expected: true,
		},
		{
			arg1:     []int{5, 1, 22, 25, 6, -1, 8, 10},
			arg2:     []int{1, 6, 10},
			expected: true,
		},
		{
			arg1:     []int{5, 1, 22, 25, 6, -1, 8, 10},
			arg2:     []int{0},
			expected: false,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - Valid subsequence should return expected output", ix), func(t *testing.T) {
			output := validSubsequence(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
