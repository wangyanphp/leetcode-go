package main

import (
	"fmt"
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {

	countMap := make(map[int]int)
	uniqueNums := make([]int, 0)

	for _, num := range nums {
		if _, ok := countMap[num]; !ok {
			countMap[num] = 0
			uniqueNums = append(uniqueNums, num)
		}
		countMap[num] += 1
	}

	sort.Ints(uniqueNums)

	for _, firstNum := range uniqueNums {

		firstNumNum := countMap[firstNum]

		if firstNumNum == 0 {
			continue
		}

		for i:=0; i< k; i++ {
			if countMap[firstNum+i] < firstNumNum {
				return false
			}
			countMap[firstNum+i] -= firstNumNum
		}
	}

	return true

}


func main() {
	nums := []int{1,2,3,3,4,4,5,6}
	k := 4

	fmt.Println(isPossibleDivide(nums, k))
}
