package balancedbinarytrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Bottom up approach
// Time complexity - O(n^2)
// Space complexity - O(n^2)
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func isBalancedTreeBottomUp(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := depth(root.Left)
	right := depth(root.Right)

	return abs(left-right) <= 1 && isBalancedTreeBottomUp(root.Left) && isBalancedTreeBottomUp(root.Right)
}

// Top down approach
// Time complexity - O(n)
// Space complexity - O(n)
func depthHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := depthHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := depthHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if (abs(leftHeight - rightHeight)) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func isBalancedTreeTopDown(root *TreeNode) bool {
	return depthHeight(root) != -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
