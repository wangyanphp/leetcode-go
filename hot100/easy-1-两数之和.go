package main

// 这里的关键是 reverse[num] = i 这句话放到哪里
// 如果放到for循环的开始处, 就没法处理[3,3], 6这种case
// 如果放到for循环的末尾处, 就可以处理
func twoSum(nums []int, target int) []int {

	reverse := make(map[int]int)
	for i, num := range nums {
		// bad
		// reverse[num] = i
		another := target - num
		if j, ok := reverse[another]; ok {
			if i > j {
				return []int{j, i}
			}else if i < j {
				return []int{i, j}
			}
		}
		reverse[num] = i
	}

	return nil
}
