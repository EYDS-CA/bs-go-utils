package api

import (
	"encoding/json"
	"net/http"

	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
)

// SnakeRequest - API request for the game
type SnakeRequest struct {
	Game  apiEntity.Game  `json:"game"`
	Turn  int             `json:"turn"`
	Board apiEntity.Board `json:"board"`
	You   apiEntity.Snake `json:"you"`
}

// DecodeSnakeRequest - turn http request into a
func DecodeSnakeRequest(req *http.Request, decoded *SnakeRequest) error {
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
}
