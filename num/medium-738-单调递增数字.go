package main

import (
	"fmt"
	"strconv"
)

var result int

func monotoneIncreasingDigits(N int) int {

	result = 0
	strN := strconv.Itoa(N)
	prefix := ""
	minV := '/' // 小于0
	__monotoneIncreasingDigits(strN, prefix, uint8(minV))
	return result
}

func __monotoneIncreasingDigits(strN string, prefix string, minV uint8) {
	dim := len(strN)

	if dim == 0 {
		cad, _ := strconv.Atoi(prefix)
		result = max(cad, result)
		return
	}

	if strN[0] == '0' {
		return
	}

	// 减1的情况
	if strN[0]-1 >= minV {
		tmpPrefix := prefix
		tmpPrefix += string(strN[0] - 1)

		for i := 0; i < len(strN)-1; i++ {
			tmpPrefix += string('9')
		}
		cad, _ := strconv.Atoi(tmpPrefix)
		result = max(cad, result)
	}

	// 不减1的情况
	// 在min - strN[0] 的最大值
	if minV > strN[0] {
		return
	}
	__monotoneIncreasingDigits(strN[1:], prefix+string(strN[0]), strN[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 算法2: 数字的基本规律
// 从数字一开始找递增的顺序, 找到第一个不满足顺序的, -1

func monotoneIncreasingDigits2(N int) int {
	chars := []uint8(strconv.Itoa(N))
	idx := 0

	for i := 0; i < len(chars)-1; i++ {

		if chars[i] < chars[i+1] {
			idx = i + 1
		}

		if chars[i] > chars[i+1] {
			chars[idx] = chars[idx] - 1
			for j := idx + 1; j < len(chars); j++ {
				chars[j] = '9'
			}
			break
		}
	}

	result, _ := strconv.Atoi(string(chars))
	return result
}

func main() {
	r := monotoneIncreasingDigits2(668841)
	fmt.Println(r)
}
