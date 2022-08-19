package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeInfo struct {
	Diameter int
	Height   int
}

func getTreeInfo(tree *TreeNode) TreeInfo {
	if tree == nil {
		return TreeInfo{0, 0}
	}

	leftTreeInfo := getTreeInfo(tree.Left)
	rightTreeInfo := getTreeInfo(tree.Right)

	largestHeightSoFar := leftTreeInfo.Height + rightTreeInfo.Height
	maxDiameterSoFar := max(leftTreeInfo.Diameter, rightTreeInfo.Diameter)
	currentDiameter := max(largestHeightSoFar, maxDiameterSoFar)
	currentHeight := 1 + max(leftTreeInfo.Height, rightTreeInfo.Height)

	return TreeInfo{currentDiameter, currentHeight}
}

func diameterOfBinaryTree(root *TreeNode) int {
	return getTreeInfo(root).Diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
