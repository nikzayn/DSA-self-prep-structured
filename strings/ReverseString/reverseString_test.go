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

func TestReverseString(t *testing.T) {
	testCases := []addTest{
		{
			str:      "nikhil",
			expected: "lihkin",
		},
		{
			str:      "sweet delhi",
			expected: "ihled teews",
		},
		{
			str:      "level",
			expected: "level",
		},
		{
			str:      "racecar",
			expected: "racecar",
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - reverse string should return expected output", ix), func(t *testing.T) {
			output := reverseString(tc.str)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
