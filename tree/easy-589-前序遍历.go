package tree

var result = make([]int, 0)

func preorder(root *Node) []int {
	result = make([]int, 0)
	preorderTraversal(root)
	return result
}

func preorderTraversal(root *Node) {
	// 非递归
	if root == nil {
		return
	}
	stack := make([]*Node, 0)

	stack = append(stack, root)

	for len(stack) != 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			result = append(result, node.Val)
		}else {
			continue
		}

		if node.Children == nil {
			continue
		}

		for i:= len(root.Children)-1; i >= 0; i-- {
			if root.Children[i] != nil {
				stack = append(stack, root.Children[i])
			}
		}
	}
}
