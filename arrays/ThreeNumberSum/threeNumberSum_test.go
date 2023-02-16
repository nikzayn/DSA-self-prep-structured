package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	arg2     int
	expected [][]int
}

func TestThreeNumberSum(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{12, 3, 1, 2, -6, 5, -8, 6},
			arg2:     0,
			expected: [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}},
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - FirstNonRepeatingCharacter should return expected output", ix), func(t *testing.T) {
			output := ThreeNumberSum(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
