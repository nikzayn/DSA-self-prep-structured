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

func TestContainsDuplicate(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			arg1:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			arg1:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - ContainsDuplicate should return expected output", ix), func(t *testing.T) {
			output := containsDuplicatesEff(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
