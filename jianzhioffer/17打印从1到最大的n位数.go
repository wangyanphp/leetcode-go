package main

import (
	"strconv"
	"strings"
)

var res = make([]string, 0)
var N int
var nums []string
// 这里的注意点是：执行之前变量初始化；迭代的时候注意全0的情况
func printNumbers(n int) []int {
	res = make([]string, 0)
	nums = make([]string, n)
	N = n
	dfs3(0, 0)

	ret := make([]int, 0)
	for _, s := range res {
		num, _ := strconv.Atoi(s)
		ret = append(ret, num)
	}

	return ret
}

func dfs3(x, prevZero int) {
	if x == N {
		// warning
		if prevZero != N {
			res = append(res, strings.Join(nums[prevZero:], ""))
		}
		return
	}

	for i := 0; i< 10; i++ {
		nums[x] = strconv.Itoa(i)
		if i == 0  && prevZero == x {
			dfs3(x+1, prevZero+1)
		}else {
			dfs3(x+1, prevZero)
		}
	}
}