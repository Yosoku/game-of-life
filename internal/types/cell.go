package types

import (
	"github.com/veandco/go-sdl2/sdl"
)

// The show is terrible, but orange and black look good on an oled panel
var AliveColor = sdl.Color{255, 95, 31, 255} // Orange
var DeadColor = sdl.Color{0, 0, 0, 0}        // Black

type Cell struct {
	X, Y  int
	Alive bool
}

func NewDeadCell(x, y int) *Cell {
	return &Cell{x, y, false}
}
func NewAliveCell(x, y int) *Cell {
	return &Cell{x, y, true}
}

func (c *Cell) Render(renderer *Renderer, w, h int32) {
	// Hacky shit to get a border
	// Render a larger black rect to act as a border
	borderRect := sdl.Rect{int32(c.X) * w, int32(c.Y) * h, w, h}
	actualRect := sdl.Rect{int32(c.X)*w + 1, int32(c.Y)*h + 1, w - 1, h - 1}
	var colour sdl.Color
	if c.Alive {
		colour = AliveColor
	} else {
		colour = DeadColor
	}
	pixel := sdl.MapRGBA(renderer.Surface.Format, colour.R, colour.G, colour.B, colour.A)
	renderer.Surface.FillRect(&borderRect, 0)
	renderer.Surface.FillRect(&actualRect, pixel)
}

func (c *Cell) updateAlive(neighbors []*Cell) {
	aliveNeighbors := getAliveCount(neighbors)
	c.Alive = aliveNeighbors == 2 || aliveNeighbors == 3
}

func (c *Cell) updateDead(neighbors []*Cell) {
	aliveNeighbors := getAliveCount(neighbors)
	c.Alive = aliveNeighbors == 3
}

func (c *Cell) Update(neighbors []*Cell) {
	if c.Alive {
		c.updateAlive(neighbors)
		return
	}
	c.updateDead(neighbors)
}

func getAliveCount(cells []*Cell) int {
	alive := 0
	for _, c := range cells {
		if c.Alive {
			alive++
		}
	}
	return alive
}
