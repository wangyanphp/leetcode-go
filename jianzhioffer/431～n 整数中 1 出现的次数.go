package main


func countDigitOne(n int) int {
	high := n/10
	low := 0
	cur := n%10
	digit := 1

	res := 0

	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		}else if cur == 1 {
			res += high*digit + low+ 1
		}else {
			res += (high+1)*digit
		}

		high, cur, low = high/10, high%10, low+cur*digit
		digit *=10
	}

	return res
}