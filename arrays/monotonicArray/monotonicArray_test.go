package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	expected bool
}

func TestMonotonicArray(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001},
			expected: true,
		},
		{
			arg1:     []int{1, 2, 0},
			expected: false,
		},
		{
			arg1:     []int{},
			expected: true,
		},
		{
			arg1:     []int{1},
			expected: true,
		},
		{
			arg1:     []int{-1, -5, 10},
			expected: false,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - move elements to end should return expected output", idx), func(t *testing.T) {
			output := isMontonicArray(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
