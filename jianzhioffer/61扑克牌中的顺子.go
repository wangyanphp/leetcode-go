package main

import "sort"

func isStraight(nums []int) bool {

	sort.Ints(nums)

	zeroNum := 0
	start := 0

	for i := 0; i< 5; i++ {
		if nums[i] == 0 {
			zeroNum++
			start++
			continue
		}
		break
	}

	for i := start+1; i < 5; i++ {
		delta := nums[i] - nums[i-1]
		if delta == 0 {
			return false
		}
		if  delta == 1 {
			continue
		}
		if delta -1  > zeroNum {
			return false
		}
		zeroNum -= delta-1
	}

	return true

}
