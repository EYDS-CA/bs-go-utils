package board

import "github.com/FreshworksStudio/bs-go-utils/api"

const (
	UP      = "up"
	DOWN    = "down"
	LEFT    = "left"
	RIGHT   = "right"
	NO_MOVE = "no_move"
)

// BoardManager for the board object
type BoardManager struct {
	GameBoard *Board
	Req       *api.SnakeRequest
	OurHead   api.Point
}

// Fill the board in based on JSON from request
func initializeBoard(req *api.SnakeRequest) *BoardManager {
	bm := new(BoardManager)
	bm.Req = req
	bm.GameBoard = createBoard(req.Board.Width, req.Board.Height)
	bm.addFood(req.Board.Food)
	bm.OurHead = bm.addSnakes(req.Board.Snakes, req.You.ID)

	return bm
}

// Add the food from the JSON
func (bm BoardManager) addFood(foodPoints []api.Point) {
	for _, element := range foodPoints {
		bm.GameBoard.insert(element, food())
	}
}

func (bm BoardManager) avgSnakeLength() float64 {
	avg := 0.0
	for _, snake := range bm.Req.Snakes {
		avg += float64(len(snake.Body))
	}
	return avg / float64(len(bm.Req.Snakes))
}

// Add our snake and the opposing snakes - with heuristic tiles
func (bm BoardManager) addSnakes(snakePoint []api.Snake, you string) api.Point {
	// Add each snake body segment
	for _, snake := range snakePoint {
		for _, snakeBody := range snake.Body {
			bm.GameBoard.insert(snakeBody, obstacle())
		}
	}

	ourHead := api.Point{}

	for _, snake := range snakePoint {
		if snake.ID == you {
			bm.GameBoard.insert(snake.Head(), snakeHead())
			ourHead = snake.Head()
		} else {
			if snake.Length > 2 {
				if distance(snake.Head(), bm.Req.You.Head()) < 5 && snake.Length >= bm.Req.You.Length {
					potential := getKillIncentive(getDirection(snake.Body[1], snake.Head()), snake.Head())
					for k, p := range potential {
						if (bm.GameBoard.tileInBounds(p)) && bm.GameBoard.getTile(p).EntityType != SNAKEHEAD {
							bm.GameBoard.insert(p, obstacle())
							if pointInSet(p, bm.Req.Food) && len(bm.Req.Food) > 1 {
								bm.Req.Food = append(bm.Req.Food[:k], bm.Req.Food[k+1:]...)
							}
						}
					}
				}
			} else {
				if snake.Length > 2 {
					potential := getKillIncentive(getDirection(snake.Body[1], snake.Head()), snake.Head())
					bm.GameBoard.insert(snake.Head(), obstacle())
					bm.Req.Food = append(bm.Req.Food, snake.Head())
					for _, p := range potential {
						if (bm.GameBoard.tileInBounds(p)) && bm.GameBoard.getTile(p).EntityType == EMPTY {
							bm.GameBoard.insert(p, food())
							if !pointInSet(p, bm.Req.Food) {
								bm.Req.Food = append(bm.Req.Food, p)
							}
						}
					}
				}
			}
		}

		if snake.Health != 100 && bm.Req.Turn > 5 {
			bm.GameBoard.insert(snake.Tail(), empty())
		}
	}

	return ourHead
}

type BestFoodResult struct {
	Differential int
	Food         Point
}

// Find the best food, the one we are closest
// to compared to all other snakes
func (bm BoardManager) findBestFood() BestFoodResult {
	best := make(map[Point]Point)
	differential := make(map[Point]int) // how much closer the person is than all other snakes
	for _, food := range bm.Req.Food {
		if distance(food, bm.OurHead) < bm.Req.You.Health {
			for _, snake := range bm.Req.Snakes {
				_, exists := best[food]
				if exists == true {
					if distance(best[food], food) > distance(snake.Head(), food) && (best[food] != food) {
						differential[food] = distance(best[food], food) - distance(snake.Head(), food)
						best[food] = snake.Head()
					}
				} else {
					best[food] = snake.Head()
					differential[food] = 15
				}
			}
		}
	}

	bestFood := BestFoodResult{0, Point{-1, -1}}
	for food := range best {
		if best[food] == bm.OurHead {
			if bestFood.Food.X == -1 {
				bestFood.Food = food
				bestFood.Differential = differential[food]
			} else {
				if distance(bestFood.Food, bm.OurHead) > distance(food, bm.OurHead) {
					bestFood.Food = food
					bestFood.Differential = differential[food]
				}
			}
		}
	}

	return bestFood
}
