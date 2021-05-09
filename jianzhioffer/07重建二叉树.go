package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var preorder, inorder []int
var dict = map[int]int{}
// 主要是采用分治法, 并且依赖于一个条件：树中的元素都不相同
func buildTree(pre []int, in []int) *TreeNode {
	preorder = pre
	inorder = in
	for i, v := range in {
		dict[v] =i
	}
	return rec(0, 0, len(in)-1)
}

//
func rec(rootI, leftI, rightI int) *TreeNode {

	if leftI > rightI {
		return nil
	}

	inorderRootI := dict[preorder[rootI]]

	root := &TreeNode{Val: preorder[rootI]}

	// 左子树
	leftRootI := rootI+1
	leftLeftI := leftI
	leftRightI := inorderRootI-1

	root.Left = rec(leftRootI, leftLeftI, leftRightI)

	// 右子树
	rightRootI := rootI+(leftRightI-leftLeftI)+2
	rightLeftI := inorderRootI+1
	rightRightI := rightI

	root.Right = rec(rightRootI, rightLeftI, rightRightI)

	return root
}

//func main()  {
//
//	pre := []int{3,9,20,15,7}
//	in := []int{9,3,15,20,7}
//
//	root := buildTree(pre, in)
//
//	_ = root
//
//}