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

func TestFirstDuplicateValue(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{2, 1, 5, 3, 3, 2, 4},
			expected: 2,
		},
		{
			arg1:     []int{2, 1, 5, 3, 3, 2, 4},
			expected: 3,
		},
		{
			arg1:     []int{1, 1, 2, 3, 3, 2, 2},
			expected: 1,
		},
		{
			arg1:     []int{},
			expected: -1,
		},
		{
			arg1:     []int{23, 21, 22, 5, 3, 13, 11, 16, 5, 11, 9, 14, 23, 3, 2, 2, 5, 11, 6, 11, 23, 8, 1},
			expected: 5,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - firstDuplicateValue should return expected output", ix), func(t *testing.T) {
			output := firstDuplicateValue(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
