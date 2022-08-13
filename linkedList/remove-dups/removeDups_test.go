package removeDups

import "testing"

func BranchSumsTest(t *testing.T) {
	got := BranchSums(&LinkedList{Val: 1, Next: &LinkedList{}})
	want := &LinkedList{}

}
