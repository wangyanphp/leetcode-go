package main

import (
	"math"
)

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	lo, hi :=  0, len(nums)-1

	for;lo <=hi; {
		if target - nums[lo] == nums[hi] {
			return []int{nums[lo], nums[hi]}
		}else if target - nums[lo] > nums[hi] {
			lo++
		}else {
			hi--
		}
	}

	return nil



}


// waning: 这里的主要是边界条件：i <= hi
func findContinuousSequence(target int) [][]int {

	res := make([][]int, 0)

	hi := target>>1

	for i := 1; i<=hi; i++ {
		j := int((math.Sqrt(float64(1-4*i+4*i*i+8*target))-1)/2)

		if j*j+j+i-i*i-2*target != 0 {
			continue
		}

		if i >j {
			continue
		}

		tmp := make([]int, j-i+1)

		for index, k := 0, i; k <= j; k++ {
			tmp[index] = k
			index++
		}
		res =append(res, tmp)
	}
	return res

}

func reverseWords(s string) string {

	return ""

}
