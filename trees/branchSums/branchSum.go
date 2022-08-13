package BranchSums

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

//Recursive DFS approach
//Time complexity - O(n)
//Space Complexity - O(n)
func BranchSums(root *BinaryTree) []int {
	// Array list initialized
	sums := []int{}
	calculateBranches(root, 0, sums)
	return sums
}

func calculateBranches(root *BinaryTree, currentVal int, sums []int) {
	if root == nil {
		return
	}
	runningSum := currentVal + root.Val
	if root.Left == nil && root.Right == nil {
		sums = append(sums, runningSum)
	}
	calculateBranches(root.Left, runningSum, sums)
	calculateBranches(root.Right, runningSum, sums)
}
