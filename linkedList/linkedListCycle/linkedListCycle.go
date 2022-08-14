package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// HashMap approach
func LinkedListCycleUsingHashmaps(head *ListNode) bool {
	cycleMap := make(map[*ListNode]struct{})

	for head != nil {
		if _, val := cycleMap[head]; val {
			return true
		}
		cycleMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// Two pointer approach
func LinkedListCycleTwoApproach(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head

	for slow != fast {
		if fast != nil || fast.Next != nil {
			return false
		}
		fast = head.Next.Next
		slow = head.Next
	}
	return true
}
