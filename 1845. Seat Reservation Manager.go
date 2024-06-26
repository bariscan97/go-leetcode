package main

import (
	"container/heap"
)

type SeatManager struct {
	last int
	pq   IntHeap
}

func Constructor(n int) SeatManager {
	return SeatManager{
		last: 0,
		pq:   make(IntHeap, 0),
	}
}

func (this *SeatManager) Reserve() int {
	if len(this.pq) == 0 {
		this.last++
		return this.last
	}
	return heap.Pop(&this.pq).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	if seatNumber == this.last {
		this.last--
	} else {
		heap.Push(&this.pq, seatNumber)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
