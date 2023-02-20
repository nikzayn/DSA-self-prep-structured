package strings

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	str      string
	expected int
}

func TestMaximumConsecutiveRepeating(t *testing.T) {
	testCases := []addTest{
		{
			str:      "AAaaabbbbaa",
			expected: 4,
		},
		{
			str:      "rtwqwaqqqqqqqqgf",
			expected: 8,
		},
		{
			str:      "a",
			expected: 0,
		},
		{
			str:      "ab",
			expected: 1,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - maximumConsecutiveRepeating should return expected output", ix), func(t *testing.T) {
			output := maximumConsecutiveRepeating(tc.str)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
