package tree

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxInt(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func AbsInt(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
