package main

import (
	"fmt"
	"strconv"
)

func minNumber(nums []int) string {

	D := 0
	for _, num := range nums {
		d := getNumDigital(num)

		if d > D {
			D = d
		}
		digits[num] = d
	}

	pow10[0] = 1
	for i := 1; i <= D; i++ {
		pow10[i] = 10*pow10[i-1]
	}

	tmp = make([]int, len(nums))

	radixsort(nums, 0, len(nums)-1, D)

	res := ""

	for _, num := range nums {
		res += strconv.Itoa(num)
	}

	return res
}

func getNumDigital(n int) int {

	if n == 0 {
		return 1
	}
	res := 0
	for; n > 0; {
		res += 1
		n /= 10
	}

	return res
}

var pow10 = make(map[int]int)
var digits = make(map[int]int)
var tmp []int

var count = make([]int, 10)
// 对[lo,hi]排序
func radixsort(nums []int, lo, hi int, D int) {

	for d := D; d >= 1; d-- {
		// 先清空
		for i :=0; i< 10; i++ {
			count[i] =0
		}

		for i := lo; i <= hi; i++ {
			count[getV(nums[i],d)]++
		}

		for i := 1; i< 10; i++ {
			count[i] = count[i-1]+count[i]
		}

		for i := hi; i>=lo; i-- {
			k := getV(nums[i], d)
			tmp[lo+count[k]-1] = nums[i]
			count[k]--
		}

		for i:= lo; i<= hi; i++ {
			nums[i] = tmp[i]
		}
	}
}

func getV(n, d int) int {
	// 返回第一位试试
	if n < pow10[d-1] {
		return (n/(pow10[digits[n]-1]))%10
	}

	return (n/(pow10[digits[n]-d]))%10
}


func main() {
	nums := []int{3,30,34,5,9}

	fmt.Println(minNumber(nums))
	fmt.Println(nums)

}
