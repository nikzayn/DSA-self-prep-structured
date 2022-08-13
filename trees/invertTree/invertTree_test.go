package trees

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
}

type ans struct {
	ans []int
}

func InvertTreeTest(t *testing.T) {
	qs := []question{

		{
			para{[]int{}},
			ans{[]int{}},
		},

		{
			para{[]int{1}},
			ans{[]int{1}},
		},

		{
			para{[]int{4, 2, 7, 1, 3, 6, 9}},
			ans{[]int{4, 7, 2, 9, 6, 3, 1}},
		},
	}

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2Preorder(invertTree(root)))
	}
	fmt.Printf("\n\n\n")
}
