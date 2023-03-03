package strings

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	str      string
	expected string
}

func TestLongestPalindromicSubstring(t *testing.T) {
	testCases := []addTest{
		{
			str:      "abaxyzzyxf",
			expected: "xyzzyx",
		},
		{
			str:      "a",
			expected: "a",
		},
		{
			str:      "it's highnoon",
			expected: "noon",
		},
		{
			str:      "abcdefgfedcbazzzzzzzzzzzzzzzzzzzz",
			expected: "zzzzzzzzzzzzzzzzzzzz",
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - longestPalindromicSubstring should return expected output", ix), func(t *testing.T) {
			output := longestPalindromicSubstring(tc.str)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
