package main

import (
	"container/heap"
	"fmt"
)

type heapOfInteger []int

func (intHeap heapOfInteger) Len() int {
	//TODO implement me
	panic("implement me")
}

func (intHeap heapOfInteger) Less(i, j int) bool {
	//TODO implement me
	panic("implement me")
}

func (intHeap heapOfInteger) GetLen() int {
	return len(intHeap)
}

func (intHeap heapOfInteger) CheckLess(i, j int) bool {
	return intHeap[i] < intHeap[j]
}

func (intHeap heapOfInteger) Swap(i, j int) {
	intHeap[i], intHeap[j] = intHeap[j], intHeap[i]
}

func (intHeap *heapOfInteger) Push(heapIntf interface{}) {
	*intHeap = append(*intHeap, heapIntf.(int))
}

func (intHeap *heapOfInteger) Pop() interface{} {
	var n int
	var x1 int
	var previous heapOfInteger = *intHeap
	n = len(previous)
	x1 = previous[n-1]
	*intHeap = previous[0 : n-1]
	return x1
}

func main() {
	var integerHeap *heapOfInteger = &heapOfInteger{1, 4, 5}
	heap.Init(integerHeap)
	heap.Push(integerHeap, 2)
	fmt.Printf("minimum: %d\n", (*integerHeap)[0])
	for integerHeap.GetLen() > 0 {
		fmt.Printf("%d \n", heap.Pop(integerHeap))
	}
}
