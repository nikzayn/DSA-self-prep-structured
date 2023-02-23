package FastAndSlowPointers

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     int
	expected bool
}

func TestHappyNumber(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     2147483646,
			expected: false,
		},
		{
			arg1:     1,
			expected: true,
		},
		{
			arg1:     19,
			expected: true,
		},
		{
			arg1:     8,
			expected: false,
		},
		{
			arg1:     7,
			expected: true,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - HappyNumber should return expected output", ix), func(t *testing.T) {
			output := HappyNumber(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
