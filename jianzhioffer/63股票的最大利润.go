package main

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minV := INT_MAX
	res := 0

	for i:= 0; i< len(prices); i++  {
		if prices[i] < minV {
			minV = prices[i]
		}else {
			res = max(prices[i]-minV, res)
		}
	}

	return res

}

func max(i, j int) int {
	if i>j {
		return i
	}
	return j
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

