package board

const (
	FOOD      = 1
	EMPTY     = 2
	SNAKEHEAD = 3
	OBSTACLE  = 4
	INVALID   = 5
)

type Entity struct {
	Display    string
	Dangerous  bool
	EntityType int
}

func GetEntity(t int) Entity {
	switch t {
	case FOOD:
		return Food()
	case EMPTY:
		return Empty()
	case SNAKEHEAD:
		return SnakeHead()
	case OBSTACLE:
		return Obstacle()
	case INVALID:
		return Invalid()
	default:
		return Invalid()
	}
}

func Empty() Entity {
	e := Entity{" ", false, EMPTY}
	return e
}

func Food() Entity {
	f := Entity{"f", false, FOOD}
	return f
}

func SnakeHead() Entity {
	h := Entity{"h", true, SNAKEHEAD}
	return h
}

func Obstacle() Entity {
	o := Entity{"o", true, OBSTACLE}
	return o
}

func Invalid() Entity {
	o := Entity{"X", true, INVALID}
	return o
}
