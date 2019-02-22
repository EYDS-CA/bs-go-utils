package game

import (
	"github.com/FreshworksStudio/bs-go-utils/api"
	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/lib"
)

// Manager for the board object
type Manager struct {
	GameBoard *Board
	Req       *api.SnakeRequest
	OurHead   apiEntity.Coord
}

// InitializeBoard - Create a manager with filled board from request
func InitializeBoard(req *api.SnakeRequest) *Manager {
	manager := new(Manager)
	manager.Req = req
	manager.GameBoard = CreateBoard(req.Board.Width, req.Board.Height)
	manager.AddFood(req.Board.Food)
	manager.AddSnakes(req.Board.Snakes, req.You.ID)
	manager.OurHead = req.You.Body[0]

	return manager
}

// AddFood - inserts food into the gameboard
func (m Manager) AddFood(food []apiEntity.Coord) {
	for _, coord := range food {
		m.GameBoard.Insert(Food(), coord)
	}
}

// AddSnakes - insert Snakes into the game board
func (m Manager) AddSnakes(snakePoint []apiEntity.Snake, you string) {

	// Add a body segment for each snake
	// Add our snake head for our head
	// Add a snake head for the opposing snakes
	for _, snake := range snakePoint {
		for index, snakeBody := range snake.Body {
			if index == 0 && snake.ID == you {
				m.GameBoard.Insert(OurSnakeHead(snake.ID), snakeBody)
			} else if index == 0 {
				m.GameBoard.Insert(SnakeHead(snake.ID), snakeBody)
			} else {
				m.GameBoard.Insert(Obstacle(snake.ID), snakeBody)
			}

			// If our health is 100, treat our tail as an obstacle
			// In the case that our head is our tail, ignore
			if index == len(snake.Body)-1 && index != 0 && snake.ID == you && lib.Distance(snake.Body[0], snake.Body[len(snake.Body)-1]) == 1 {
				m.GameBoard.Insert(Empty(), snakeBody)
			}

			if len(snake.Body) > 1 {
				if snake.ID == you && snake.Body[0] != snake.Body[1] {
					m.GameBoard.Insert(Obstacle(you), snake.Body[1])
				}
			}
		}
	}

}
