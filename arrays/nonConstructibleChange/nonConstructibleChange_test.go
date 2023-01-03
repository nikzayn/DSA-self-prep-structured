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

func TestNonConstructibleChange(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{1, 2, 5},
			expected: 4,
		},
		{
			arg1:     []int{5, 7, 1, 1, 2, 3, 22},
			expected: 20,
		},
		{
			arg1:     []int{1, 1, 1, 1, 1, 1},
			expected: 7,
		},
		{
			arg1:     []int{109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4},
			expected: 87,
		},
		{
			arg1:     []int{1, 1},
			expected: 3,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - sorted squares should return expected output", idx), func(t *testing.T) {
			output := nonConstructibleChange(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
