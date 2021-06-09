package main

// 这里有几个重要的点：
// 1. m,n 代表的是长度，这样的话，状态转移数组是m+1*n+1
// 2. 想好状态转移方程
// 3. dp[i][j] 依赖于 dp[i-xx][j]所以，嵌套方式为i，j
func isMatch(s string, p string) bool {

	// dp[i][j] 代表长度为i和j
	m, n := len(s)+1, len(p)+1

	dp := make([][]bool, m)
	for i := 0; i< m; i++ {
		dp[i] = make([]bool, n)
	}

	dp[0][0] = true
	for i :=1; i< m; i++ {
		dp[i][0] = false
	}

	for i := 0; i < m; i++ {
		for j := 1; j < n; j ++ {
			if p[j-1] != '*' {
				if i >= 1 {
					dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
				}
			}else {
				dp[i][j] = (j >= 2 && dp[i][j-2]) || (i >= 1 && j>= 2 && dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}

	return dp[m-1][n-1]
}
