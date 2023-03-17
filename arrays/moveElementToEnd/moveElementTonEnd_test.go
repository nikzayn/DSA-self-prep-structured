package arrays

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

func TestMoveElementsToLast(t *testing.T) {
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
			arg1:     []int{3, 3, 3, 3, 3},
			arg2:     3,
			expected: []int{3, 3, 3, 3, 3},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - move elements to end should return expected output", idx), func(t *testing.T) {
			output := moveElementToEndEff(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
