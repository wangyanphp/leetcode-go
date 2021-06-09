package main


func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if A ==  nil {
		return false
	}
	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)

}

func isSub(A, B *TreeNode) bool {

	if B == nil {
		return true
	}
	if A ==  nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
}
