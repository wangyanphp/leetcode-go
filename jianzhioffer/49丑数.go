package main

func nthUglyNumber(n int) int {

	dp := make([]int, n+1)

	dp[1] = 1
	i2, i3, i5 := 1,1,1

	for i := 2; i <= n; i++ {

		a2, a3, a5 := dp[i2]*2, dp[i3]*3, dp[i5]*5

		dp[i] = min(min(a2, a3),a5)
		if dp[i] == a2 {
			i2++
		}

		if dp[i] == a3 {
			i3++
		}
		if dp[i] == a5 {
			i5++
		}

	}

	return dp[n]


}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


