package strings

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     string
	arg2     string
	expected int
}

func TestOccurenceOfFirstIndexString(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     "sadbutsad",
			arg2:     "sad",
			expected: 0,
		},
		{
			arg1:     "leetcode",
			arg2:     "leeto",
			expected: -1,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - occurence of first index string should return expected output", ix), func(t *testing.T) {
			output := OccurenceOfFirstIndexString(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
