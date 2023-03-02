package validParenthesis

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     string
	expected bool
}

func TestValidParentheses(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     "()[]{}",
			expected: true,
		},
		{
			arg1:     "((()))",
			expected: true,
		},
		{
			arg1:     "[]{)))",
			expected: false,
		},
		{
			arg1:     "[[(({{}}))]]",
			expected: true,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - valid parentheses should return expected output", idx), func(t *testing.T) {
			output := validParenthesis(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
