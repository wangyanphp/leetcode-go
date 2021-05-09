package main

// 这里的关键是搞成一个二叉树
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	// m -> 行, n -> 列
	m, n := len(matrix), len(matrix[0])

	for i,j := 0, n-1;i<=m-1 && j>=0;{

		//// 边界条件
		//if i > m-1 || j < 0 {
		//	return false
		//}
		if matrix[i][j] == target {
			return true
		}else if target > matrix[i][j] {
			i++
		}else {
			j--
		}
	}

	return false
}

