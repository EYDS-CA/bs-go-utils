package apiEntity

// Coord - Destructured Coord object for request
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Snake - Destructured Snake object for request
type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []Coord `json:"body"`
}

// Board - Destructured Board object for request
type Board struct {
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Food   []Coord `json:"food"`
	Snakes []Snake `json:"snakes"`
}

// Game - Destructured Game object for request
type Game struct {
	ID string `json:"id"`
}
