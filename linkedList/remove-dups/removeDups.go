package removeDups

type LinkedList struct {
	Val  int
	Next *LinkedList
}

// Approach is to iterate to the list, remove the existing by updating currentNode to nextNode
// and vice versa
// Time complexity - O(n) - because we are searching a particular node
// Space complexity - O(1) - since updating value to the node
func RemoveDups(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList
	for currentNode != nil {
		nextNode := currentNode.Next
		for nextNode != nil && currentNode.Val == nextNode.Val {
			nextNode = nextNode.Next
		}
		// Update the currentNode next to nextNode
		currentNode.Next = nextNode
		// Update the currentNode to nextNode Next
		currentNode = nextNode.Next
	}
	return linkedList
}
