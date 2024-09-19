package types

var Left = Direction{-1, 0}
var Right = Direction{1, 0}
var Up = Direction{0, 1}
var Down = Direction{0, -1}

type Direction struct {
	x, y int
}

func Add(a, b Direction) Direction {
	return Direction{a.x + b.x, a.y + b.y}
}

var AllDirections = []Direction{Left, Right, Up, Down, Add(Up, Left), Add(Up, Right), Add(Down, Left), Add(Down, Right)}
