package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 这里的关键是: 处理好next指针, 如何递进
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	result := &ListNode{}

	head := result

	addTo := 0
	for ; ; {

		if l1 == nil || l2 == nil {
			break
		}
		result.Next = &ListNode{}
		result = result.Next

		v1 := l1.Val
		v2 := l2.Val

		r := v1 + v2 + addTo

		if r >= 10 {
			result.Val = r - 10
			addTo = 1
		} else {
			result.Val = r
			addTo = 0
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 == nil {
		if addTo != 0 {
			result.Next = &ListNode{}
			result = result.Next
			result.Val = addTo
		}
		return head.Next
	}

	var l *ListNode
	if l1 == nil {
		l = l2
	} else {
		l = l1
	}

	for ; ; {

		if l == nil {
			break
		}

		result.Next = &ListNode{}
		result = result.Next

		v1 := l.Val

		r := v1 + addTo

		if r >= 10 {
			result.Val = r - 10
			addTo = 1
		} else {
			result.Val = r
			addTo = 0
		}

		l = l.Next
	}

	if addTo != 0 {
		result.Next = &ListNode{}
		result = result.Next
		result.Val = addTo
	}
	return head.Next

}

// func main() {
// 	l1 := &ListNode{
// 		Val:  2,
// 		Next: &ListNode{
// 			Val:  4,
// 			Next: &ListNode{
// 				Val:  3,
// 			},
// 		},
// 	}
//
// 	l2 := &ListNode{
// 		Val:  5,
// 		Next: &ListNode{
// 			Val:  6,
// 			Next: &ListNode{
// 				Val:  4,
// 			},
// 		},
// 	}
//
// 	r := addTwoNumbers(l1, l2)
// 	fmt.Println(r)
// }
