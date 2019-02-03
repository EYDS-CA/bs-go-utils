package api

// StartResponse - Returns the color to be used during the game
type StartResponse struct {
	Color string `json:"color,omitempty"`
}

// MoveResponse - Returns the move to be taken for a turn
type MoveResponse struct {
	Move string `json:"move"`
}

// EmptyResponse - Generic Empty response for Ping and end
type EmptyResponse struct{}
