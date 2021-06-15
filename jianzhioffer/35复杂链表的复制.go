package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}

//
func copyRandomList1(head *Node) *Node {
	maps := make(map[*Node]*Node)

	for tmp := head; tmp != nil; tmp = tmp.Next {
		maps[tmp] = &Node{
			Val:    tmp.Val,
		}
	}

	for tmp := head; tmp != nil; tmp = tmp.Next {
		maps[tmp].Next = maps[tmp.Next]
		maps[tmp].Random = maps[tmp.Random]
	}

	return maps[head]
}


func copyRandomList(head *Node) *Node {

	// 处理好边界
	if head == nil {
		return nil
	}

	for tmp := head; tmp != nil; {
		newNode := &Node{
			Val:    tmp.Val,
			Next:   nil,
		}
		oldNext := tmp.Next
		tmp.Next = newNode
		newNode.Next =  oldNext
		tmp = oldNext
	}



	for tmp := head; tmp != nil; tmp = tmp.Next.Next{
		if tmp.Random != nil {
			tmp.Next.Random = tmp.Random.Next
		}
	}

	newHead := head.Next

	var tmp, tmpNew *Node
	for tmp, tmpNew = head, head.Next; tmpNew.Next != nil;  {

		tmp.Next = tmp.Next.Next
		tmpNew.Next = tmpNew.Next.Next
		tmp = tmp.Next
		tmpNew = tmpNew.Next
	}

	// warning: 还原好现场
	tmp.Next = nil

	return newHead

}
