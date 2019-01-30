package game

import (
	"container/heap"
	"fmt"

	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/lib"
)

// Path - a list of Coords
type Path = []*apiEntity.Coord

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

// FindPath - Use A* to find Path from start Coord to end Coord
func (m Manager) FindPath(start apiEntity.Coord, end apiEntity.Coord) Path {

	// closedSet - Coordinates that have already been explored
	closedSet := make(map[apiEntity.Coord]bool)

	// openSet - Coordinates to be explored, start at
	openSet := make(CoordHeap, 0)
	heap.Init(&openSet)
	heap.Push(&openSet, &CoordNode{Coord: start, Value: 0})

	// cameFrom - Keep track of how a coordinate got from one to other
	// use when the path needs to be reconstructed
	cameFrom := make(map[apiEntity.Coord]apiEntity.Coord)

	// gScore - How many steps it took to get to that tile as a path
	gScore := make(map[apiEntity.Coord]float32)

	// fScore - How far away the tile is from the goal (Manhatten)
	fScore := make(map[apiEntity.Coord]float32)

	// Deffault g & f score to be a really high number
	for i := 0; i < m.GameBoard.Width; i++ {
		for j := 0; j < m.GameBoard.Height; j++ {
			gScore[apiEntity.Coord{X: i, Y: j}] = 10000.0
			fScore[apiEntity.Coord{X: i, Y: j}] = float32(lib.Distance(apiEntity.Coord{X: i, Y: j}, end))
		}
	}

	// It takes 0 steps to get to the start
	gScore[start] = 0

	for len(openSet) > 0 {

		// Pull the current closest node off the heap
		min := heap.Pop(&openSet).(*CoordNode)
		fmt.Printf("AT COORDINATE %+v\n", min.Coord)

		// If the current closest node is the goal, reconstruct the path
		if min.Coord.X == end.X {
			fmt.Printf("GOT TO THE GOAL")
		}

		// Add the current coordinate to the already explored tiles
		closedSet[min.Coord] = true

		// Get the current coordinates adjacent tiles so that
		// the tile can be explored later
		neighbours := m.GameBoard.GetValidTiles(min.Coord)

		// Explore the neighbours
		for _, neighbor := range neighbours {

			// We have already explored the neighbor, ignore it
			_, inClosedSet := closedSet[neighbor]
			if inClosedSet {
				continue
			}

			// The current score for the neighbor is the number
			// of steps it took to ge there + direct distance
			tentativeGScore := gScore[min.Coord] + float32(lib.Distance(min.Coord, neighbor))

			// if !pointInSet(n, openSet) {
			// 	openSet = append(openSet, n)
			// } else if tentativeGScore >= gScore[n] {
			// 	continue
			// }

			// If we have come across a quicker path to a Coordinate, record it
			if tentativeGScore <= gScore[neighbor] {
				heap.Push(&openSet, &CoordNode{Coord: neighbor, Value: int(tentativeGScore)})
			} else {
				continue
			}

			cameFrom[neighbor] = min.Coord
			gScore[neighbor] = tentativeGScore
			fScore[neighbor] = tentativeGScore + float32(lib.Distance(neighbor, min.Coord))
		}

	}

	fmt.Printf("%+v\n", cameFrom)
	return make(Path, 0)
}
