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

func TestMinimumRemoveParenthesess(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     "(((abc)(to)((q)()(",
			expected: "(abc)(to)(q)()",
		},
		{
			arg1:     ")((ab)c))(ac(op)t(())(",
			expected: ")((ab)c))ac(op)t(())",
		},
		{
			arg1:     "((((sop(fc(aq((z)(a)())(((",
			expected: "sopfcaq((z)(a)())",
		},
		{
			arg1:     ")()r(s(t(q(v)(w(x)y)z())((()(",
			expected: "()rst(q(v)(w(x)y)z())()",
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - minimumRemoveParentheses should return expected output", idx), func(t *testing.T) {
			output := minimumRemoveParentheses(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
