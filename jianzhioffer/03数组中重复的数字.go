package main

// 这里的一个点是：数组长度为n，数字全部都在n-1中，只有所有数字都出现且唯一出现一次的时候，才不会重复
// 怎么实现这一点
// 对于索引i，如果nums[i] = i，那么ok
// nums[i] < i 则有重复
// nums[i] 与 nums[nums[i]]互换，直到nums[i] <= i

func findRepeatNumber(nums []int) int {

	l := len(nums)
	for i :=0 ; i<l; {

		if nums[i] == i {
			i++
			continue
		}

		if nums[i] < i {
			return nums[i]
		}

		for ;nums[i] > i; {
			// warning: 没有这个条件[1,1,1]就没法结束
			if nums[i]==nums[nums[i]]{
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return 0
}

//func main()  {
//
//	nums := []int{2, 3, 1, 0, 2, 5, 3}
//	r := findRepeatNumber(nums)
//	fmt.Println(r)
//
//}