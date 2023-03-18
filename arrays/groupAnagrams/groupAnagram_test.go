package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type addTest struct {
	arg1     []string
	expected [][]string
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []addTest{
		{
			arg1:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			arg1:     []string{""},
			expected: [][]string{{""}},
		},
		{
			arg1:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for ix, tc := range testCases {
		t.Run(fmt.Sprintf("test %d - GroupAnagrams should return expected output", ix), func(t *testing.T) {
			output := groupAnagrams(tc.arg1)

			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", output, tc.expected)
			}
		})
	}
}
