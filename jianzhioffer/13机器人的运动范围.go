package main

// 一个深度or广度优先遍历
// 这里的问题在于防止重复遍历
// 这里的优化点是：
// 1. 下标值的计算是递推的
// 2. 虽然可以前进上下左右，但是实际上从（0，0）开始只需要右边和下面
// 3. 注意边界条件别冲突
var marked [][]bool

func movingCount(m int, n int, k int) int {
	marked = make([][]bool, m)
	for i := 0; i<m; i++ {
		marked[i] = make([]bool, n)
	}

	return dfs2(0, 0, m, n, k, 0, 0)
}

func dfs2(i,j,m,n, k, si, sj int) int {

	if i>=m || i<0 || j >= n || j < 0 || si+sj > k {
		return 0
	}

	if marked[i][j] {
		return 0
	}

	marked[i][j] = true

	siadd1 := si
	if (i+1) % 10 != 0 {
		siadd1 +=1
	}else {
		siadd1 -= 8
	}

	sjadd1 := sj
	if (j+1) %10 != 0 {
		sjadd1 += 1
	}else {
		sjadd1 -= 8
	}

	return 1 + dfs2(i + 1, j, m,n, k, siadd1, sj) + dfs2(i, j + 1, m, n, k, si, sjadd1)
}
