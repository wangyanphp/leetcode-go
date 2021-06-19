package main

func lastRemaining(n int, m int) int {

	dp := make([]int, n)

	dp[0] = 0
	// warning: 这里的坑点在于对i+1取模
	for i := 1; i< n; i++ {
		dp[i] = (m+dp[i-1])%(i+1)
	}

	return dp[n-1]

}
