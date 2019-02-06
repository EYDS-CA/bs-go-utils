package game

import (
	"fmt"
	"strings"

	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
)

// Board - Object to hold game data
type Board struct {
	Width  int
	Height int
	Grid   [][]Entity
}

// CreateBoard - returns an width x height board of empties
func CreateBoard(width int, height int) *Board {
	b := new(Board)
	b.Width = width
	b.Height = height
	grid := make([][]Entity, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]Entity, width)
		for j := 0; j < width; j++ {
			grid[i][j] = Empty()
		}
	}

	b.Grid = grid
	return b
}

// XInBounds - return if an X coordinate is within the grid
func (b Board) XInBounds(xpos int) bool {
	return (0 <= xpos && xpos < b.Width)
}

// YInBounds - return if an X coordinate is within the grid
func (b Board) YInBounds(ypos int) bool {
	return (0 <= ypos && ypos < b.Height)
}

// GetTile returns the entity at Coord(x,y) or Invalid
func (b Board) GetTile(c apiEntity.Coord) Entity {
	if b.TileInBounds(c) {
		return b.Grid[c.Y][c.X]
	}
	return Invalid()
}

// TileInBounds - Returns if a coordinate(x,y) is within the grid
func (b Board) TileInBounds(c apiEntity.Coord) bool {
	return (b.XInBounds(c.X) && b.YInBounds(c.Y))
}

// Insert - Insert entity e at coordinate c
func (b Board) Insert(e Entity, c apiEntity.Coord) {
	if b.XInBounds(c.X) && b.YInBounds(c.Y) {
		b.Grid[c.Y][c.X] = e
	}
}

// GetValidTiles - for a coordinate c
// return all Coordinates which are not obstacles
// or are not out of bounds
func (b Board) GetValidTiles(c apiEntity.Coord) []apiEntity.Coord {
	validTiles := make([]apiEntity.Coord, 0)
	potential := []apiEntity.Coord{
		apiEntity.Coord{X: c.X - 1, Y: c.Y},
		apiEntity.Coord{X: c.X + 1, Y: c.Y},
		apiEntity.Coord{X: c.X, Y: c.Y - 1},
		apiEntity.Coord{X: c.X, Y: c.Y + 1},
	}
	for i, c := range potential {
		if b.TileInBounds(c) && b.GetTile(c).EntityType != OBSTACLE {
			validTiles = append(validTiles, potential[i])
		}
	}
	return validTiles
}

// Show - Pretty Prints the board and all its entities
func (b Board) Show() {
	rowDivider := strings.Repeat(" ---", b.Width)
	println(rowDivider)
	for i := 0; i < b.Height; i++ {
		print("| ")
		for j := 0; j < b.Width; j++ {
			fmt.Printf("%s | ", b.Grid[i][j].Display)
		}
		println("\n" + rowDivider)
	}
}
