package main

func search(nums []int, target int) int {

	lo, hi := 0, len(nums)-1
	left := bsearch(nums, lo, hi, target, true)
	if left < 0 {
		return 0
	}
	right := bsearch(nums, lo, hi, target, false)

	return right-left+1

}

func bsearch(nums []int, lo,hi int, target int, leftFlag bool) int {

	if lo > hi {
		return -1
	}

	length := hi-lo+1
	if length == 1 {
		if nums[lo] == target {
			return lo
		}else {
			return -1
		}
	}

	mid := lo + length/2 -1

	if target > nums[mid] {
		return bsearch(nums, mid+1, hi, target, leftFlag)
	}else if target < nums[mid] {
		return bsearch(nums, lo, mid-1, target, leftFlag)
	}

	if leftFlag {
		if mid -1 >= lo && nums[mid-1] == target {
			return bsearch(nums, lo, mid-1, target, leftFlag)
		}else {
			return mid
		}
	}else {
		if mid +1 <=hi && nums[mid+1] == target {
			return bsearch(nums, mid+1,hi, target, leftFlag)
		}else {
			return mid
		}
	}
}




