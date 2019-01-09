package board

import (
	"github.com/FreshworksStudio/bs-go-utils/api"
)

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
func InitializeBoard(req *api.SnakeRequest) *BoardManager {
	bm := new(BoardManager)
	bm.Req = req
	bm.GameBoard = CreateBoard(req.Board.Width, req.Board.Height)
	bm.addFood(req.Board.Food)
	bm.OurHead = bm.addSnakes(req.Board.Snakes, req.You.ID)

	return bm
}

func (bm BoardManager) avgSnakeLength() float64 {
	avg := 0.0
	for _, snake := range bm.Req.Board.Snakes {
		avg += float64(len(snake.Body))
	}
	return avg / float64(len(bm.Req.Board.Snakes))
}

// Add the food from the JSON
func (bm BoardManager) addFood(foodPoints []api.Point) {
	for _, element := range foodPoints {
		bm.GameBoard.insert(element, Food())
	}
}

// Add our snake and the opposing snakes - with heuristic tiles
func (bm BoardManager) addSnakes(snakePoints []api.Snake, you string) api.Point {
	// Add each snake body segment
	for _, snake := range snakePoints {
		for _, snakeBody := range snake.Body {
			bm.GameBoard.insert(snakeBody, Obstacle())
		}
	}

	ourHead := api.Point{}

	for _, snake := range snakePoints {
		if snake.ID == you {
			bm.GameBoard.insert(snake.Head(), SnakeHead())
			ourHead = snake.Head()
		} else {
			if len(snake.Body) > 2 {
				if Distance(snake.Head(), bm.Req.You.Head()) < 5 && len(snake.Body) >= len(bm.Req.You.Body) {
					potential := GetKillIncentive(GetDirection(snake.Body[1], snake.Head()), snake.Head())
					for k, p := range potential {
						if (bm.GameBoard.tileInBounds(p)) && bm.GameBoard.getTile(p).EntityType != SNAKEHEAD {
							bm.GameBoard.insert(p, Obstacle())
							if PointInSet(p, bm.Req.Board.Food) && len(bm.Req.Board.Food) > 1 {
								bm.Req.Board.Food = append(bm.Req.Board.Food[:k], bm.Req.Board.Food[k+1:]...)
							}
						}
					}
				}
			} else {
				if len(snake.Body) > 2 {
					potential := GetKillIncentive(GetDirection(snake.Body[1], snake.Head()), snake.Head())
					bm.GameBoard.insert(snake.Head(), Obstacle())
					bm.Req.Board.Food = append(bm.Req.Board.Food, snake.Head())
					for _, p := range potential {
						if (bm.GameBoard.tileInBounds(p)) && bm.GameBoard.getTile(p).EntityType == EMPTY {
							bm.GameBoard.insert(p, Food())
							if !PointInSet(p, bm.Req.Board.Food) {
								bm.Req.Board.Food = append(bm.Req.Board.Food, p)
							}
						}
					}
				}
			}
		}

		if snake.Health != 100 && bm.Req.Turn > 5 {
			bm.GameBoard.insert(snake.Tail(), Empty())
		}
	}

	return ourHead
}
