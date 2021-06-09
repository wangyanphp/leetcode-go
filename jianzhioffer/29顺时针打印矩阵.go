package main


// warning: 得处理好边界条件
func spiralOrder(matrix [][]int) []int {

	res := make([]int, 0)

	m := len(matrix)
	if m == 0 {
		return res
	}
	n := len(matrix[0])

	startI, endI := 0, m-1
	startJ, endJ := 0, n-1

	for; startI < endI &&  startJ < endJ; {
		for k := startJ; k <= endJ; k++ {
			res = append(res, matrix[startI][k])
		}


		for k := startI+1; k <= endI; k++ {
			res = append(res, matrix[k][endJ])
		}

		for k := endJ-1; k>= startI; k-- {
			res = append(res, matrix[endI][k])
		}

		for k := endI-1; k >= startI+1; k-- {
			res = append(res, matrix[k][startJ])
		}
		startI++
		endI--
		startJ++
		endJ--
	}

	if startI == endI &&  startJ == endJ {
		res = append(res, matrix[startI][startJ])
	}else if startI == endI && startJ < endJ {
		for k := startJ; k <= endJ; k++ {
			res = append(res, matrix[startI][k])
		}
	}else if startI < endI && startJ == endJ {
		for k := startI; k <= endI; k++ {
			res = append(res, matrix[k][endJ])
		}
	}

	return res





}
