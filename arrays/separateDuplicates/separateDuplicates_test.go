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

func TestSeparateDuplicates(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{2, 3, 3, 3, 6, 9, 9},
			expected: 4,
		},
		{
			arg1:     []int{2, 2, 2, 11},
			expected: 2,
		},
		{
			arg1:     []int{3, 2, 3, 6, 3, 10, 9, 3},
			expected: 5,
		},
		{
			arg1:     []int{0},
			expected: 1,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d -  should return expected output", idx), func(t *testing.T) {
			output := SeparateDuplicatesNaive(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
