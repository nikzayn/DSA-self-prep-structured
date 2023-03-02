package stacks

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     string
	expected string
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     "g",
			expected: "g",
		},
		{
			arg1:     "ggaabcdeb",
			expected: "bcdeb",
		},
		{
			arg1:     "abbddaccaaabcd",
			expected: "abcd",
		},
		{
			arg1:     "aabbccdd",
			expected: "",
		},
		{
			arg1:     "aannkwwwkkkwna",
			expected: "kwkwna",
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - valid parentheses should return expected output", idx), func(t *testing.T) {
			output := removeDuplicates(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
