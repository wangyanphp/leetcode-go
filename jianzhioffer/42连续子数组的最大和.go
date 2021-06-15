package main

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

func maxSubArray(nums []int) int {

	res, max := INT_MIN, INT_MIN
	boolNeg := true
	// 如果全是负数，选择一个最大的负数
	for _, num := range nums {
		if num > 0 {
			boolNeg = false
			break
		}
		if num > res {
			res =  num
		}
	}

	if boolNeg {
		return res
	}


	res = nums[0]
	max = nums[0]

	for _, num := range nums[1:] {
		res += num

		if res > max {
			max = res
		}
		if res < 0 {
			res = 0
		}
	}

	return max

}