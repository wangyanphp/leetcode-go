package main

// 这里的几个坑点：1. d的pow计算；2. n是索引，不是长度，然后还有0存在，所以边界条件需要弄好
// 3. 对于d位数字，第一个数字需要特殊处理

func findNthDigit(n int) int {

	// 先判断是几位数
	d := 1
	if n <=9 {
		return n
	}

	for r := 9 ;; {
		if n > r*d {
			n -= r*d
			d++
			r = r*10
			continue
		}
		break
	}
	// 第一个数字的处理比较特殊
	if n <= d {
		if n == 1 {
			return 1
		}else {
			return 0
		}
	}


	// 在d位数，偏移量n
	pre := n/d
	mod := n%d

	num := pre-1+pow10(d-1)
	// 就是num的最后一位
	if mod == 0 {
		return num % 10
	}else {
		return ((num+1)/(pow10(d-mod)))%10
	}


}

func pow10(n int) int {
	if n == 0 {
		return 1
	}

	x := 10
	res := 1
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