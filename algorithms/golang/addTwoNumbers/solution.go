package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
		}
		if l2 != nil {
			carry += l2.Val
		}

		cur.Next = &ListNode{
			Val:  carry % 10,
		}
		carry = carry / 10
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		cur.Next = &ListNode{
			Val:  carry,
		}
	}

	return dummy.Next
}

func helper(l1 *ListNode, l2 *ListNode, plus bool) {
	if l2 != nil {
		l1.Val = l1.Val + l2.Val
	}
	if plus {
		l1.Val += 1
		plus = false
	}
	if l1.Val >= 10 {
		l1.Val = l1.Val % 10
		plus = true
	}

	if l1.Next == nil {
		if l2 == nil || l2.Next == nil {
			return
		}
		l1.Next = l2.Next
		helper(l1.Next, nil, plus)
	}
	if l2 != nil {
		l2 = l2.Next
	}
	helper(l1.Next, l2, plus)
}