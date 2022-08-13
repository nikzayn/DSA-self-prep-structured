package invertTree

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

type TreeNodes struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive DFS + Swapping Approach
// Time Complexity - O(n)
// Space Complexity - O(h)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	swapLeftAndRight(root)
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTreeWithoutSwapFunc(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root

}

func swapLeftAndRight(root *TreeNode) {
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}
