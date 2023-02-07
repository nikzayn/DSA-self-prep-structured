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

func TestFirstNonRepeatingCharacter(t *testing.T) {
	testCases := []addTest{
		{
			str:      "abcdcaf",
			expected: 1,
		},
		{
			str:      "faadabcbbebdf",
			expected: 6,
		},
		{
			str:      "a",
			expected: 0,
		},
		{
			str:      "ab",
			expected: 0,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - FirstNonRepeatingCharacter should return expected output", ix), func(t *testing.T) {
			output := NaiveFirstNonRepeatingCharacter(tc.str)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
