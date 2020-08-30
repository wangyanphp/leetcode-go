package tree

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		stack: make([]*TreeNode, 0),
	}
	for root != nil {
		it.stack = append(it.stack, root)
		root = root.Left
	}
	return it
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {

	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	root := val.Right
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
	return val.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}