package main

var path []int
var res [][]int
// 需要找出每一个path，path就得当参数
// 支持负数
func pathSum(root *TreeNode, target int) [][]int {
	res  = make([][]int, 0)
	path = make([]int, 0)
	dfs5(root, target)
	return res
}

func dfs5(root *TreeNode, target int)  {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Val == target && root.Left == nil && root.Right == nil {
		tmp := make([]int,len(path))
		copy(tmp,path)
		res = append(res, tmp)
	}

	dfs5(root.Left, target-root.Val)
	dfs5(root.Right, target-root.Val)
	path = path[:len(path)-1]
}