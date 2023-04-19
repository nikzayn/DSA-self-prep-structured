package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AddTest struct {
	arg      []int
	expected int
}

func TestMissingNum(t *testing.T) {
	assert := assert.New(t)
	tests := []AddTest{
		{
			arg:      []int{1, 2, 3, 4, 5, 7},
			expected: 6,
		},
		{
			arg:      []int{1, 2, 3, 4, 6, 7, 8, 9, 10},
			expected: 5,
		},
	}

	for _, test := range tests {
		assert.Equal(findMissingNum(test.arg), test.expected)
	}
}

func TestMissingNumBenchmark(b *testing.B) {
	arr := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		findMissingNum(arr)
	}
}
