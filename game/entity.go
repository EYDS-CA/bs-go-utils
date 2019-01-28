package game

/**
 * FOOD: A food which can be eaten
 * EMPTY: A tile with nothing in it
 * SNAKEHEAD: The head of a snake, like an obstacle
 * OBSTACLE: A snake body, do not touch
 * INVALID: An invalid entry, used in place of null
 */
const (
	FOOD      = 1
	EMPTY     = 2
	SNAKEHEAD = 3
	OBSTACLE  = 4
	INVALID   = 5
)

// Entity - An entity at a Coord(x, y)
type Entity struct {
	Display    string
	Dangerous  bool
	EntityType int
	SnakeID    string
}

// Empty - create an Empty
func Empty() Entity {
	return Entity{" ", false, EMPTY, ""}
}

// Food -  create a Food
func Food() Entity {
	return Entity{"F", false, FOOD, ""}
}

// SnakeHead - create a SnakeHead
func SnakeHead(snakeID string) Entity {
	return Entity{"H", true, SNAKEHEAD, snakeID}
}

// OurSnakeHead - Add our SnakeHead
func OurSnakeHead(snakeID string) Entity {
	return Entity{"Y", true, SNAKEHEAD, snakeID}
}

// Obstacle - create an Obstacle
func Obstacle(snakeID string) Entity {
	return Entity{"O", true, OBSTACLE, snakeID}
}

// Invalid - return an invalid entry
func Invalid() Entity {
	return Entity{"X", true, INVALID, ""}
}
