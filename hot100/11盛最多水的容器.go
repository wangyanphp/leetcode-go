package main

import "fmt"

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	i, j := 0, len(height)-1

	left := height[i]
	right := height[j]
	res := (j-i)* min(left, right)

	for i < j {

		if left < right {
			for i < j {
				if height[i] <= left {
					i++
					continue
				}
				left = height[i]
				break
			}
		}else {
			for i < j {
				if height[j] <= right {
					j--
					continue
				}
				right = height[j]
				break
			}
		}
		tmp := (j-i)* min(left, right)
		if tmp > res {
			res = tmp
		}
	}

	return res



}


func min(i, j int) int {
	if i<j {
		return i
	}
	return j
}

func main() {
	nums := []int{1,8,6,2,5,4,8,3,7}

	v := maxArea(nums)
	fmt.Println(v)
}
