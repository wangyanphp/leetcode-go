package main

type ListNode struct {
   Val int
     Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {

	prev := head
	cur := head

	if head.Val == val {
		return head.Next
	}

	for cur = head.Next; cur != nil; prev, cur = prev.Next, cur.Next{
		if cur.Val == val {
			prev.Next = cur.Next
			return head
		}
	}

	return head

}
