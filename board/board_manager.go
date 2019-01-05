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
	for _, snake := range bm.Req.Board.Snakes {
		avg += float64(len(snake.Body))
	}
	return avg / float64(len(bm.Req.Board.Snakes))
}

// Add our snake and the opposing snakes - with heuristic tiles
func (bm BoardManager) addSnakes(snakePoints []api.Snake, you string) api.Point {
	// Add each snake body segment
	for _, snake := range snakePoints {
		for _, snakeBody := range snake.Body {
			bm.GameBoard.insert(snakeBody, obstacle())
		}
	}

	ourHead := api.Point{}

	for _, snake := range snakePoints {
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

	return api.Point{0, 0}
}
