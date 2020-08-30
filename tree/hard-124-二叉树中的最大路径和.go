package tree

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	intMin := ^(int(^uint(0) >> 1))

	val := intMin

	maxPathSumWithValue(root, &val)
	return val

}

// 主要在于: 区分递归部分和非递归部分
func maxPathSumWithValue(root *TreeNode, val *int) int {
	if root == nil {
		return 0
	}

	leftMax := maxPathSumWithValue(root.Left, val)
	rightMax := maxPathSumWithValue(root.Right, val)

	lmr := root.Val + MaxInt(leftMax, 0) + MaxInt(rightMax, 0)

	ret := root.Val + MaxInt(0, MaxInt(leftMax, rightMax))
	*val = MaxInt(*val, MaxInt(lmr, ret))
	return ret
}

