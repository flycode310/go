package main

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	if k <=0 || k > len(nums) {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	h := &myHeap{}
	heap.Init(h)
	for _,value := range nums {
		if h.Len() < k {
			heap.Push(h, value)
			continue
		}
		if value > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, value)
		}
	}
	return (*h)[0]
}


//构建堆

type myHeap []int

//实现对应接口

func (h myHeap)Len()int{
	return len(h)
}

func (h myHeap)Less(i,j int)bool{
	return h[i]<h[j]
}

func (h myHeap)Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}

func (h *myHeap)Push(x interface{}){
	*h = append(*h,x.(int))
}

func (h *myHeap)Pop()interface{}{
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

