package main

type heapOfInteger []int

func (intHeap heapOfInteger) GetLen() int {
	return len(intHeap)
}

func (intHeap heapOfInteger) CheckLess(i, j int) bool {
	return intHeap[i] < intHeap[j]
}

func (intHeap heapOfInteger) Swap(i, j int) {
	intHeap[i], intHeap[j] = intHeap[j], intHeap[i]
}
