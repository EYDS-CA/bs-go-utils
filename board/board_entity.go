package board

const (
	FOOD      = 1
	EMPTY     = 2
	SNAKEHEAD = 3
	OBSTACLE  = 4
	INVALID   = 5
)

type entity struct {
	Display    string
	Dangerous  bool
	EntityType int
}

func getEntity(t int) entity {
	switch t {
	case FOOD:
		return food()
	case EMPTY:
		return empty()
	case SNAKEHEAD:
		return snakeHead()
	case OBSTACLE:
		return obstacle()
	case INVALID:
		return invalid()
	default:
		return invalid()
	}
}

func empty() entity {
	e := entity{" ", false, EMPTY}
	return e
}

func food() entity {
	f := entity{"f", false, FOOD}
	return f
}

func snakeHead() entity {
	h := entity{"h", true, SNAKEHEAD}
	return h
}

func obstacle() entity {
	o := entity{"o", true, OBSTACLE}
	return o
}

func invalid() entity {
	o := entity{"X", true, INVALID}
	return o
}
