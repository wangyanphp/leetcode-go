package main

func verifyPostorder(postorder []int) bool {

	if len(postorder) == 0 || len(postorder) == 1 {
		return true
	}

	root := postorder[len(postorder)-1]

	// 左子树全都小于root，右子树全都大于root
	mid := 0
	for ; mid < len(postorder)-1; mid++ {
		if postorder[mid] > root {
			break
		}
	}

	for i := mid; i < len(postorder)-1; i++ {
		if postorder[i] < root {
			return false
		}
	}

	return verifyPostorder(postorder[:mid]) && verifyPostorder(postorder[mid:len(postorder)-1])
}
