package main

import (
	"container/heap"
	"fmt"
)

// Returns the Absolute Value of an integer
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PointNode struct {
	Value int
	Point point
	Index int
}

// A Heap for performing A* Search on nodes
type PointHeap []*PointNode

// Len returns the number of elements in the heap
func (h PointHeap) Len() int {
	return len(h)
}

// Less compares the values of the node[i] and node[j] for sorting
func (h PointHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

// Pop Removes the smallest element from the heap and returns it
func (h *PointHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	item := old[n-1]
	item.Index = -1
	*h = old[0 : n-1]
	return item
}

// Push element onto the heap
func (h *PointHeap) Push(x interface{}) {
	n := h.Len()
	item := x.(*PointNode)
	item.Index = n
	*h = append(*h, item)
}

// Swap the values of items at index i,j
func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func main() {
	listNodePoints := []*PointNode{
		{Point: point{X: 1, Y: 2}, Value: 30},
		{Point: point{X: 3, Y: 5}, Value: 4},
		{Point: point{X: 2, Y: 2}, Value: 50},
		{Point: point{X: 4, Y: 4}, Value: 900},
	}
	pointHeap := make(PointHeap, len(listNodePoints))
	for i, item := range listNodePoints {
		pointHeap[i] = item
		pointHeap[i].Index = i
	}
	heap.Init(&pointHeap)

	node := heap.Pop(&pointHeap).(*PointNode)
	fmt.Printf("Value: %s, Point (X: %d, Y: %d) \n", node.Value, node.Point.X, node.Point.Y)

	heap.Push(&pointHeap, &PointNode{Point: point{X: 2, Y: 2}, Value: 23})

	node = heap.Pop(&pointHeap).(*PointNode)
	fmt.Printf("Value: %s, Point (X: %d, Y: %d): \n", node.Value, node.Point.X, node.Point.Y)

	node = heap.Pop(&pointHeap).(*PointNode)
	fmt.Printf("Value: %s, Point (X: %d, Y: %d): \n", node.Value, node.Point.X, node.Point.Y)

	node = heap.Pop(&pointHeap).(*PointNode)
	fmt.Printf("Value: %s, Point (X: %d, Y: %d): \n", node.Value, node.Point.X, node.Point.Y)
}
