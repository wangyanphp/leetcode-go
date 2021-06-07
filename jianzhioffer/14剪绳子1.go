package main

import "math"

// 这里很明显的是用数学知识，分解成尽可能2和3。
// 边界条件： n=2 => 1*1 n=3 => 2*1 n=4 => 2*2

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a - 1)) * 4)
	}

	return int(math.Pow(3, float64(a)) * 2)
}
