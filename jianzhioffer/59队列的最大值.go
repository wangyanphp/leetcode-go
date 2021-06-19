package main

import (
	"fmt"
)

type MaxQueue struct {
	q []int
	maxQ []int
}


func Constructor() MaxQueue {
	return MaxQueue{
		q:    make([]int, 0),
		maxQ: make([]int, 0),
	}
}


func (this *MaxQueue) Max_value() int {

	if len(this.maxQ) > 0 {
		return this.maxQ[0]
	}
	return -1

}


func (this *MaxQueue) Push_back(value int)  {

	this.q = append(this.q, value)

	for len(this.maxQ) != 0 && value > this.maxQ[len(this.maxQ)-1]{
		this.maxQ = this.maxQ[:len(this.maxQ)-1]
	}
	this.maxQ = append(this.maxQ,value)
}


func (this *MaxQueue) Pop_front() int {

	ret := -1
	if len(this.q)  == 0 {
		return ret
	}

	ret = this.q[0]
	this.q = this.q[1:]
	if ret == this.maxQ[0] {
		this.maxQ = this.maxQ[1:]
	}
	return ret

}



func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	maxIndexQ := make([]int, 0)
	res := make([]int, 0)

	for i := 0; i < k; i++ {
		for len(maxIndexQ) != 0 && nums[i] > nums[maxIndexQ[len(maxIndexQ)-1]]{
			maxIndexQ = maxIndexQ[:len(maxIndexQ)-1]
		}
		maxIndexQ = append(maxIndexQ, i)
	}


	res = append(res, nums[maxIndexQ[0]])

	for i := k; i< len(nums); i++ {

		for len(maxIndexQ) != 0 && nums[i] > nums[maxIndexQ[len(maxIndexQ)-1]]{
			maxIndexQ = maxIndexQ[:len(maxIndexQ)-1]
		}
		maxIndexQ = append(maxIndexQ, i)

		// 删除元素
		if i-k == maxIndexQ[0] {
			maxIndexQ = maxIndexQ[1:]
		}
		res = append(res, nums[maxIndexQ[0]])
	}

	return res

}

func main()  {

	nums := []int{7,2,4}

	fmt.Println(maxSlidingWindow(nums, 2))

}
