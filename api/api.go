package api

import (
	"encoding/json"
	"net/http"
)

const (
	HEAD_BENDR     HeadType = "bendr"
	HEAD_DEAD               = "dead"
	HEAD_FANG               = "fang"
	HEAD_PIXEL              = "pixel"
	HEAD_REGULAR            = "regular"
	HEAD_SAFE               = "safe"
	HEAD_SAND_WORM          = "sand-worm"
	HEAD_SHADES             = "shades"
	HEAD_SMILE              = "smile"
	HEAD_TONGUE             = "tongue"
)

const (
	TAIL_BLOCK_BUM    TailType = "block-bum"
	TAIL_CURLED                = "curled"
	TAIL_FAT_RATTLE            = "fat-rattle"
	TAIL_FRECKLED              = "freckled"
	TAIL_PIXEL                 = "pixel"
	TAIL_REGULAR               = "regular"
	TAIL_ROUND_BUM             = "round-bum"
	TAIL_SKINNY                = "skinny"
	TAIL_SMALL_RATTLE          = "small-rattle"
)

type HeadType string

type TailType string

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PointList []Point

type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []Point `json:"body"`
}

func (snake Snake) Head() Point { return snake.Body[0] }

func (snake Snake) Tail() Point { return snake.Body[len(snake.Body)-1] }

type SnakeList []Snake

type Board struct {
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Food   []Point `json:"food"`
	Snakes []Snake `json:"snakes"`
}

type Game struct {
	ID string `json:"id"`
}

type SnakeRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

type StartResponse struct {
	Color          string   `json:"color,omitempty"`
	Name           string   `json:"name,omitempty"`
	HeadURL        string   `json:"head_url,omitempty"`
	Taunt          string   `json:"taunt,omitempty"`
	HeadType       HeadType `json:"head_type,omitempty"`
	TailType       TailType `json:"tail_type,omitempty"`
	SecondaryColor string   `json:"secondary_color,omitempty"`
}

type MoveResponse struct {
	Move string `json:"move"`
}

func DecodeSnakeRequest(req *http.Request, decoded *SnakeRequest) error {
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
}
