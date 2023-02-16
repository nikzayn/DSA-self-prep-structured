package strings

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
			arg1:     "level",
			expected: true,
		},
		{
			arg1:     "racecear",
			expected: false,
		},
		{
			arg1:     "gear",
			expected: false,
		},
		{
			arg1:     "racecar",
			expected: true,
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - FirstNonRepeatingCharacter should return expected output", ix), func(t *testing.T) {
			output1 := NaiveValidPalindrome(tc.arg1)
			output2 := EffValidPalindrome(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output1) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output1, tc.expected)
			}

			if !reflect.DeepEqual(tc.expected, output2) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output2, tc.expected)
			}
		})
	}
}
