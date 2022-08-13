package trees

import (
	"testing"
)

func BranchSumsTest(t *testing.T) {
	got := BranchSums(&BinaryTree{Val: 1, Left: &BinaryTree{}, Right: &BinaryTree{}})
	want := []int{}

}
