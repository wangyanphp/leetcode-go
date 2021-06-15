package main

import "strconv"

func translateNum(num int) int {

	str := strconv.Itoa(num)

	dp := make([]int, len(str))

	dp[0] = 1

	for i := 1; i < len(str); i++ {
		dp[i] = dp[i-1]
		if str[i-1] != '0' && (str[i-1] - '0')*10+(str[i]-'0') <= 25 {

			if i >= 2 {
				dp[i] += dp[i-2]
			}else {
				dp[i] += 1
			}

		}
	}
	return dp[len(str)-1]
}
