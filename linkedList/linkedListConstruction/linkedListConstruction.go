package linkedListConstruction

// Node - represents a node of linked list
type Node struct {
	Val  int
	Next *Node
}

// Linked list - represents a linked list
type LinkedList struct {
	head *Node
	len  *LinkedList
}

func initList() *LinkedList {
	return &LinkedList{}
}

// Crud operation in Linked list
// Insert nodes in a linked list
func (l *LinkedList) InsertBefore() {

}
