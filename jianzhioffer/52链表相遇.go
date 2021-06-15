package main

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	for ; a != nil && b != nil; a, b = a.Next, b.Next {
		if a == b {
			return a
		}
	}

	aN := 0
	bN := 0

	if a == nil && b != nil {
		for; b != nil ; b = b.Next {
			bN++
		}
	}else if a != nil && b == nil {
		for; a!=nil; a =a.Next {
			aN++
		}
	}

	a, b = headA, headB
	for i :=0; i<aN; i++ {
		a = a.Next
	}

	for i := 0; i<bN; i++ {
		b = b.Next
	}

	for ; a != nil && b != nil; a, b = a.Next, b.Next {
		if a == b {
			return a
		}
	}

	return nil

}


// warning: 不相交的话，a和b都是走过m+n长度即可
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	for ; a != b ; {

		if a == nil {
			a = headB
		}else {
			a = a.Next
		}

		if b == nil {
			b = headA
		}else {
			b = b.Next
		}
	}

	return a

}

