package main

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{
		Val: root.Val,
	}
	newRoot.Left = mirrorTree(root.Right)
	newRoot.Right = mirrorTree(root.Left)
	return newRoot
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recSym(root.Left, root.Right)

}

func recSym(L, R *TreeNode) bool {

	if L == nil && R == nil {
		return true
	}

	if L == nil || R == nil {
		return false
	}

	return L.Val == R.Val && recSym(L.Left, R.Right) && recSym(L.Right, R.Left)
}
