package main

// 广度优先
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for ; len(queue) > 0; {
		item := queue[0]
		queue = queue[1:]
		res = append(res, item.Val)
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
	}
	return res
}

// 当遍历的时候，收尾就是
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	q1 := make([]*TreeNode, 0)
	q2 := make([]*TreeNode, 0)

	q1 = append(q1, root)

	for ; len(q1)> 0 || len(q2) > 0; {
		level := make([]int, 0)
		for ; len(q1) > 0; {
			item := q1[0]
			q1 = q1[1:]
			level = append(level, item.Val)
			if item.Left != nil {
				q2 = append(q2, item.Left)
			}
			if item.Right != nil {
				q2 = append(q2, item.Right)
			}
		}
		res = append(res, level)
		q1, q2 = q2, q1
	}

	return res
}



func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	q1 := make([]*TreeNode, 0)
	q2 := make([]*TreeNode, 0)

	q1 = append(q1, root)
	reFlag := false
	for ; len(q1)> 0 || len(q2) > 0; {
		tmpLevel := make([]int, 0)
		for ; len(q1) > 0; {
			item := q1[0]
			q1 = q1[1:]
			tmpLevel = append(tmpLevel, item.Val)
			if item.Left != nil {
				q2 = append(q2, item.Left)
			}
			if item.Right != nil {
				q2 = append(q2, item.Right)
			}
		}

		level := make([]int, 0)
		if reFlag {
			for i := len(tmpLevel)-1; i >= 0; i-- {
				level = append(level, tmpLevel[i])
			}
		}else {
			level = tmpLevel
		}

		res = append(res, level)
		q1, q2 = q2, q1
		reFlag = !reFlag
	}

	return res
}
