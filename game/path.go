package game

import "github.com/FreshworksStudio/bs-go-utils/apiEntity"

type Path struct {
}

// CoordNode - Heap Node containing a Coord
type CoordNode struct {
	Value int
	Coord apiEntity.Coord
	Index int
}

// CoordHeap - Heap of Coords
type CoordHeap []*CoordNode

// Len returns the number of elements in the heap
func (h CoordHeap) Len() int {
	return len(h)
}

// Less compares the values of the node[i] and node[j] for sorting
func (h CoordHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

// Pop Removes the smallest element from the heap and returns it
func (h *CoordHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	item := old[n-1]
	item.Index = -1
	*h = old[0 : n-1]
	return item
}

// Push element onto the heap
func (h *CoordHeap) Push(x interface{}) {
	n := h.Len()
	item := x.(*CoordNode)
	item.Index = n
	*h = append(*h, item)
}

// Swap the values of items at index i,j
func (h CoordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}
