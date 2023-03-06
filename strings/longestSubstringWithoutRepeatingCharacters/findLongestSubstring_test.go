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

func TestFindLongestSubstring(t *testing.T) {
	testCases := []addTest{
		{
			str:      "abcabcbb",
			expected: 3,
		},
		{
			str:      "bbbbb",
			expected: 1,
		},
		{
			str:      "pwwkew",
			expected: 3,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - findLongestSubstring should return expected output", ix), func(t *testing.T) {
			output := findLongestSubstring(tc.str)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
