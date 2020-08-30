package tree

func isSubPath(head *ListNode, root *TreeNode) bool {

	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	return isSubPathNoSkip(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSubPathNoSkip(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	if root.Val != head.Val {
		return false
	}

	return isSubPathNoSkip(head.Next, root.Left) || isSubPathNoSkip(head.Next, root.Right)
}

