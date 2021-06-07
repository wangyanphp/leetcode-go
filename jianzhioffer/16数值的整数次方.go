package main

// 一个比较简单的想法是：变成了计算x^(n/2) 然后相乘即可
// 这里可以用上动态规划存储一下子
// 但是没有必要这么重，
// x^2^(n/2) -> 这个可能一下子计算出来
//
//func myPow(x float64, n int) float64 {
//
//	if n < 0 {
//		return 1/ myPow(x, -n)
//	}
//	if n == 0 {
//		return 1
//	}
//
//	if n & 1 == 1{
//		return x * myPow(x*x, (n-1)/2)
//	}
//
//	return myPow(x*x, (n-1)/2)
//}

// 如何把递归改成迭代？
// n -> n/2
func myPow(x float64, n int) float64 {

	if n < 0 {
		return 1/ myPow(x, -n)
	}
	if n == 0 {
		return 1
	}

	res := float64(1)
	// res -> 1
	// TODO: 这里的主要问题是设置n的条件是1，而不是0
	for ;n> 1; {
		if n & 1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	res *= x
	return res
}