package tree

// 2020.8.30 二叉树的递归和动态规划在很大程度上需要总结一下的,
// 二叉树天然分叉, 可能会自动剪纸
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)

	if leftHeight == -1 {
		return -1
	}
	rightHeight := height(root.Right)

	if rightHeight == -1 {
		return -1
	}
	if AbsInt(leftHeight - rightHeight) > 1 {
		return -1
	}
	return MaxInt(leftHeight, rightHeight) + 1
}


