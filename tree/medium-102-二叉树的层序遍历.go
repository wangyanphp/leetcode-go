package tree

func levelOrder(root *TreeNode) [][]int {

	ret := make([][]int, 0)

	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		items := make([]int, 0)

		for i := 0; i < l; i++ {
			tmp := queue[0]
			queue = queue[1:]
			items = append(items, tmp.Val)
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		ret = append(ret, items)
	}

	return ret

}


// func main() {
// 	root := &TreeNode{
// 		Val:   3,
// 		Left:  &TreeNode{
// 			Val:   9,
// 			Left:  nil,
// 			Right: nil,
// 		},
// 		Right: &TreeNode{
// 			Val:   20,
// 			Left:  &TreeNode{Val: 15},
// 			Right: &TreeNode{Val: 7},
// 		},
// 	}
// 	ret := levelOrder(root)
// 	fmt.Println(ret)
// }

