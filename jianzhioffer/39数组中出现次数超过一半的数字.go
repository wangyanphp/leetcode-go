package main

// 这里的关键是摩尔投票法：如果有占一半以上的数字，那就返回之，否则结果未定义
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	
	var res, resVote int
	
	for i := 0; i < len(nums); i++ {
		if resVote == 0 {
			resVote = 1
			res = nums[i]
			continue
		}
		
		if nums[i] == res {
			resVote ++
		}else {
			resVote --
		}
	}
	
	count := 0
	
	for i := 0; i < len(nums); i++ {
		if nums[i] == res {
			count++
		}else {
			count --
		}
	}
	if count >= 0 {
		return res
	}
	return 0

}
