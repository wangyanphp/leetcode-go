package main

import (
	"container/heap"
	"fmt"
)

// 这里的关键是 如何使用golang实现优先队列
// addNum的关键之处在于，新增数字必须在两个堆中都走一遍


type Heap []int

func (h Heap) Len() int {
	return len(h)
}
func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
// Pop 弹出最后一个元素
func (h *Heap) Pop() interface{}{
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type MinHeap struct {
	Heap
}
func (h MinHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h.Heap[i] <  h.Heap[j]
}

type MaxHeap struct {
	Heap
}
func (h MaxHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h.Heap[i] >  h.Heap[j]
}


type MedianFinder struct {
	*MinHeap
	*MaxHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{
		MinHeap: &MinHeap{Heap:make(Heap, 0)},
		MaxHeap: &MaxHeap{Heap:make(Heap, 0)},
	}
	heap.Init(m.MinHeap)
	heap.Init(m.MaxHeap)
	return m
}



func (this *MedianFinder) AddNum(num int)  {


	if this.MinHeap.Len() == this.MaxHeap.Len() {
		heap.Push(this.MinHeap, num)
		heap.Push(this.MaxHeap, heap.Pop(this.MinHeap))
	}else {
		heap.Push(this.MaxHeap, num)
		heap.Push(this.MinHeap, heap.Pop(this.MaxHeap))
	}
}


func (this *MedianFinder) FindMedian() float64 {

	if this.MaxHeap.Len() == 0 {
		return 0
	}
	if this.MinHeap.Len() != this.MaxHeap.Len() {
		return float64(this.MaxHeap.Heap[0])
	}
	tmp1 := this.MinHeap.Heap[0]
	tmp2 := this.MaxHeap.Heap[0]
	return (float64(tmp1+tmp2))/2
}