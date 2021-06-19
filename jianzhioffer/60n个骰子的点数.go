package main

// warning: 这里捋清楚关系很关键
func dicesProbability(n int) []float64 {

	dp := make([]float64, 6)
	for i := 0; i< 6; i++ {
		dp[i] = float64(1)/6
	}

	for i := 2; i <=n; i++ {
		tmp := make([]float64, i*5+1)

		// 就是这一段
		for k := 0; k < len(dp); k++ {
			for cur := 0; cur < 6; cur++ {
				tmp[k+cur] += dp[k]/6
			}
		}
		dp = tmp
	}

	return dp

}

