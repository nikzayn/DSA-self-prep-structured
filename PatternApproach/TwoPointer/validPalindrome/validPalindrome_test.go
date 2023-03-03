package TwoPointer

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     string
	expected bool
}

func TestValidPalindrome(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     "akkai",
			expected: false,
		},
		{
			arg1:     "level",
			expected: true,
		},
		{
			arg1:     "racecar",
			expected: true,
		},
		{
			arg1:     "nikhil",
			expected: false,
		},
	}
	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - Valid palindrome should return expected output", ix), func(t *testing.T) {
			output := EffValidPalindrome(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
