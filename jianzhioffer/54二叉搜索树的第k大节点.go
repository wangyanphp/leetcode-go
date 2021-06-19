package main

func kthLargest(root *TreeNode, k int) int {

	val, _ := rec(root, k)
	if val != nil {
		return *val
	}
	return 0

}

func rec(root *TreeNode, k int) (val *int, count int) {

	if root == nil {
		return nil, 0
	}

	// 找到了
	val, rightC := rec(root.Right, k)

	if val != nil {
		return val, rightC
	}

	// 没找到, root
	if rightC == k-1 {
		return &root.Val, k
	}

	val, leftC := rec(root.Left, k- rightC-1)
	if val != nil {
		return val, leftC
	}

	return nil, leftC+rightC+1
}



func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left+1
	}
	return right+1
}