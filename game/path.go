package game

import (
	"container/heap"
	"errors"
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

// Helper function - see if an Coord has already been added to the heap
func coordInSlice(c apiEntity.Coord, s CoordHeap) bool {
	for i := 0; i < len(s); i++ {
		if c.X == s[i].Coord.X && c.Y == s[i].Coord.Y {
			return true
		}
	}
	return false
}

// ReversePath - Reverses a path, useful because path is generated backwards
func ReversePath(path Path) Path {
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// ReconstructPath - given the map of Coords and where they came from, return a Path
func ReconstructPath(current CoordNode, pathMap map[apiEntity.Coord]apiEntity.Coord) Path {
	path := make(Path, 0)
	path = append(path, &current.Coord)
	fmt.Printf("%v+  %v+", pathMap[apiEntity.Coord{X: 1, Y: 3}], pathMap[apiEntity.Coord{X: 1, Y: 4}])

	finished := false
	for !finished {
		// fmt.Printf("AT %v+\n", current.Coord)
		current, exists := pathMap[current.Coord]
		if !exists {
			finished = true
		} else {
			// fmt.Printf("NOW at %v+\n", current)
			path = append(path, &current)
		}
	}

	return ReversePath(path)
}

// FindPath - Use A* to find Path from start Coord to end Coord
func (m Manager) FindPath(start apiEntity.Coord, end apiEntity.Coord) (Path, error) {
	fmt.Printf("STAR: %v+, FINISH: %v+\n", start, end)
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
		if min.Coord.X == end.X && min.Coord.Y == end.Y {
			fmt.Printf("GOT TO THE GOAL\n")
			return ReconstructPath(*min, cameFrom), nil
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
				// fmt.Printf("ALREADY SEEN NEIGHBOR %v+, SKIPPING\n", neighbor)
				continue
			}

			// The current score for the neighbor is the number
			// of steps it took to ge there + direct distance
			tentativeGScore := gScore[min.Coord] + float32(lib.Distance(min.Coord, neighbor))

			if !coordInSlice(neighbor, openSet) {
				fmt.Printf("FOUND NEW COORD, %v+, PUSHING TO OPEN\n", neighbor)
				heap.Push(&openSet, &CoordNode{Coord: neighbor, Value: int(tentativeGScore)})
			} else if tentativeGScore >= gScore[neighbor] {
				continue
			}
			cameFrom[neighbor] = min.Coord
			gScore[neighbor] = tentativeGScore
			fScore[neighbor] = tentativeGScore + float32(lib.Distance(neighbor, min.Coord))
		}
	}

	return nil, errors.New("Unable to find a path")
}
