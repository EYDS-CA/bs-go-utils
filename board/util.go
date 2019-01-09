package board

import (
	"math"

	"github.com/FreshworksStudio/bs-go-utils/api"
)

func ToStringPointer(str string) *string {
	return &str
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func PointInSet(p api.Point, s []api.Point) bool {
	for i := 0; i < len(s); i++ {
		if p.X == s[i].X && p.Y == s[i].Y {
			return true
		}
	}

	return false
}

func Distance(p1 api.Point, p2 api.Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

func ReconstructPath(current api.Point, pathMap map[api.Point]api.Point) []api.Point {
	path := make([]api.Point, 0)
	path = append(path, current)

	_, exists := pathMap[current]

	for ; exists; _, exists = pathMap[current] {
		current = pathMap[current]
		path = append(path, current)
	}

	return ReverseList(path)
}

func ProjectSnakeAlongPath(path []api.Point, snake api.Snake) []api.Point {
	if len(path) < len(snake.Body) {
		p := make([]api.Point, 0)
		p = append(p, path[:len(path)]...)
		p = append(p, snake.Body[:(len(snake.Body)-len(path))+1]...)
		return p
	} else if len(path) > len(snake.Body) {
		return path[:len(snake.Body)]
	}

	return path
}

func PathIsSafe(path []api.Point, ourSnake api.Snake, b *Board) bool {
	path = ReverseList(path)
	if len(path) < 2 {
		return false
	}

	copy := b.copy()
	for _, v := range ourSnake.Body {
		copy.insert(v, Empty())
	}

	projected := ProjectSnakeAlongPath(path, ourSnake)
	for _, p := range projected {
		copy.insert(p, Obstacle())
	}
	fakeHead := projected[0]
	fakeTail := projected[len(projected)-1]
	copy.insert(fakeHead, SnakeHead())
	copy.insert(fakeTail, Empty())

	pathToTail := ShortestPath(fakeHead, fakeTail, copy)
	if len(pathToTail) > 2 {
		return true
	}

	return false
}

func ReverseList(lst []api.Point) []api.Point {
	for i := 0; i < len(lst)/2; i++ {
		j := len(lst) - i - 1
		lst[i], lst[j] = lst[j], lst[i]
	}
	return lst
}

func GetDirection(from api.Point, to api.Point) string {
	vertical := to.Y - from.Y
	horizontal := to.X - from.X
	if vertical == 0 {
		if horizontal > 0 {
			return RIGHT
		}
		return LEFT
	}
	if vertical < 0 {
		return UP
	}
	return DOWN
}

func PairIsValidExtension(p1 api.Point, p2 api.Point, board Board, path []api.Point) bool {
	return PointIsValidExtension(p1, board, path) && PointIsValidExtension(p2, board, path)
}

func PointIsValidExtension(p api.Point, board Board, path []api.Point) bool {
	return !board.getTile(p).Dangerous && !PointInSet(p, path)
}

func ExtendPath(path []api.Point, board Board, limit int) []api.Point {
	extended := make([]api.Point, 0)
	extended = append(extended, path...)
	for i := 0; i < len(extended)-1; i++ {
		current := extended[i]
		next := extended[i+1]
		direction := GetDirection(current, next)
		if direction == RIGHT || direction == LEFT {
			currentUp := api.Point{current.X, current.Y - 1}
			currentDown := api.Point{current.X, current.Y + 1}
			nextUp := api.Point{next.X, next.Y - 1}
			nextDown := api.Point{next.X, next.Y + 1}
			if PairIsValidExtension(currentUp, nextUp, board, extended) {
				extended = append(extended[0:i+1], append([]api.Point{currentUp, nextUp}, extended[i+1:]...)...)
			} else if PairIsValidExtension(currentDown, nextDown, board, extended) {
				extended = append(extended[0:i+1], append([]api.Point{currentDown, nextDown}, extended[i+1:]...)...)
			}
		} else if direction == UP || direction == DOWN {
			currentLeft := api.Point{current.X - 1, current.Y}
			currentRight := api.Point{current.X + 1, current.Y}
			nextLeft := api.Point{next.X - 1, next.Y}
			nextRight := api.Point{next.X + 1, next.Y}
			if PairIsValidExtension(currentLeft, nextLeft, board, extended) {
				extended = append(extended[0:i+1], append([]api.Point{currentLeft, nextLeft}, extended[i+1:]...)...)
			} else if PairIsValidExtension(currentRight, nextRight, board, extended) {
				extended = append(extended[0:i+1], append([]api.Point{currentRight, nextRight}, extended[i+1:]...)...)
			}
		}
		if i == len(extended)-1 || len(extended) > limit {
			continue
		}
	}
	return extended
}

// Find the shortest path from start -> goal
func ShortestPath(start api.Point, goal api.Point, board *Board) []api.Point {
	closedSet := make([]api.Point, 0) // Tiles already explored
	openSet := make([]api.Point, 0)   // Tiles to explore
	openSet = append(openSet, start)  // Start exploring from start tile

	gScore := make(map[api.Point]float32) // Shortest path distance
	fScore := make(map[api.Point]float32) // Manhatten distance heuristic
	cameFrom := make(map[api.Point]api.Point)
	for i := 0; i < board.Width; i++ {
		for j := 0; j < board.Height; j++ {
			gScore[api.Point{i, j}] = 1000.0
			fScore[api.Point{i, j}] = 1000.0
		}
	}
	gScore[start] = 0
	fScore[start] = float32(Distance(start, goal))

	// While there are still tiles to explore
	for len(openSet) > 0 {
		// Pick the current closest based on the heuristic
		min := openSet[0]
		minIndex := 0
		for i := 0; i < len(openSet); i++ {
			if fScore[openSet[i]] < fScore[min] {
				min = openSet[i]
				minIndex = i
			}
		}
		if min.X == goal.X && min.Y == goal.Y {
			// fmt.Println("got here")
			return ReconstructPath(goal, cameFrom)
		}

		// Remove the minimum from the open set, add to closed set
		openSet[minIndex] = openSet[len(openSet)-1]
		openSet = openSet[:len(openSet)-1] // << maybe here?
		closedSet = append(closedSet, min)
		neighbours := board.getValidTiles(min)

		// Explore the neighbours
		for _, n := range neighbours {
			if PointInSet(n, closedSet) {
				continue
			}

			tentativeGScore := gScore[min] + float32(Distance(min, n))

			if !PointInSet(n, openSet) {
				openSet = append(openSet, n)
			} else if tentativeGScore >= gScore[n] {
				continue
			}

			cameFrom[n] = min
			gScore[n] = tentativeGScore

			var bonus float32
			if board.getTile(n).EntityType == EMPTY {
				bonus = -0.1
			} else {
				bonus = 0.0
			}

			fScore[n] = tentativeGScore + float32(Distance(n, min)) + bonus
		}
	}

	return nil
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func GetKillIncentive(direction string, head api.Point) []api.Point {
	switch direction {
	case UP:
		return []api.Point{
			api.Point{head.X - 1, head.Y - 1},
			api.Point{head.X, head.Y - 1},
			api.Point{head.X + 1, head.Y - 1},
		}
	case LEFT:
		return []api.Point{
			api.Point{head.X - 1, head.Y - 1},
			api.Point{head.X - 1, head.Y},
			api.Point{head.X - 1, head.Y + 1},
		}
	case DOWN:
		return []api.Point{
			api.Point{head.X - 1, head.Y + 1},
			api.Point{head.X, head.Y + 1},
			api.Point{head.X + 1, head.Y + 1},
		}
	case RIGHT:
		return []api.Point{
			api.Point{head.X + 1, head.Y - 1},
			api.Point{head.X + 1, head.Y},
			api.Point{head.X + 1, head.Y + 1},
		}
	default:
		return []api.Point{
			api.Point{head.X - 1, head.Y - 1},
			api.Point{head.X, head.Y - 1},
			api.Point{head.X + 1, head.Y - 1},
		}
	}
}
