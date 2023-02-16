package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	expected int
}

func TestSecondLargest(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{2, 3, 1, 5},
			expected: 3,
		},
		{
			arg1:     []int{2, 3, 5},
			expected: 3,
		},
		{
			arg1:     []int{5},
			expected: -1,
		},
		{
			arg1:     []int{1, 0, 0, 1, 2},
			expected: 1,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - second largest should return expected output", idx), func(t *testing.T) {
			output := SecondLargest(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
