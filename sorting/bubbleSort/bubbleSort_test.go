package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []int
	expected []int
}

func TestBubbleSort(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []int{8, 5, 2, 9, 5, 6, 3},
			expected: []int{2, 3, 5, 5, 6, 8, 9},
		},
		{
			arg1:     []int{1},
			expected: []int{1},
		},
		{
			arg1:     []int{-7, 2, 3, 8, -10, 4, -6, -10, -2, -7, 10, 5, 2, 9, -9, -5, 3, 8},
			expected: []int{-10, -7, -7, -6, -6, -5, -5, -4, -4, -4, -2, -1, 1, 3, 5, 5, 6, 8, 8, 10},
		},
		{
			arg1:     []int{427, 787, 222, 996, -359, -614, 246, 230, 107, -706, 568, 9, -246, 12, -764, -212, -484, 603, 934, -848, -646, -991, 661, -32, -348, -474, -439, -56, 507, 736, 635, -171, -215, 564, -710, 710, 565, 892, 970, -755, 55, 821, -3, -153, 240, -160, -610, -583, -27, 131},
			expected: []int{-991, -848, -764, -755, -710, -706, -646, -614, -610, -583, -484, -474, -439, -359, -348, -246, -215, -212, -171, -160, -153, -56, -32, -27, -3, 9, 12, 55, 107, 131, 222, 230, 240, 246, 427, 507, 564, 565, 568, 603, 635, 661, 710, 736, 787, 821, 892, 934, 970, 996},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - bubble sort array should return expected output", idx), func(t *testing.T) {
			output := bubbleSort(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
