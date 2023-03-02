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
			str:      "We love Go",
			expected: "Go love We",
		},
		{
			str:      "To be or not to be",
			expected: "be to not or be To",
		},
		{
			str:      "You are amazing",
			expected: "amazing are You",
		},
		{
			str:      "Hello     World",
			expected: "World Hello",
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - reverse string should return expected output", ix), func(t *testing.T) {
			output := ReverseString(tc.str)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
