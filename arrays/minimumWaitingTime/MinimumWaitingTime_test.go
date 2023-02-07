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

func TestMinimumWaitingTime(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{3, 2, 1, 2, 6},
			expected: 17,
		},
		{
			arg1:     []int{2, 1, 1, 1},
			expected: 6,
		},
		{
			arg1:     []int{1, 2, 4, 5, 2, 1},
			expected: 23,
		},
		{
			arg1:     []int{25, 30, 2, 1},
			expected: 32,
		},
		{
			arg1:     []int{7},
			expected: 2,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - MinimumWaitingTime should return expected output", ix), func(t *testing.T) {
			output := MinimumWaitingTime(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
