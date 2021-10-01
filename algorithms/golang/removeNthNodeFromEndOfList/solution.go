package removeNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}

	fast := head
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil ; {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
