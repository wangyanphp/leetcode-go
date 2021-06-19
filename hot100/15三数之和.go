package main

import (
	"sort"
)

// warning: 这里的主要问题是：如何去重并且不漏掉元素。
// 以[-1 -1 2]为例， 这是一个合法的，所以不能上来就用index 1，而应该上来先用index 0，而后跳过index 1
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	res := make([][]int, 0)

	 for i := 0; i < len(nums)-1; i++ {
		// 需要跟上次枚举的不同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		k := len(nums)-1

		for j := i+1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[j] + nums[k] > target {
				k--
			}

			if j>=k {
				break
			}

			if nums[j] + nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return res
}