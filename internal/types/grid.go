package types

type Grid struct {
	Width, Height         int
	CellWidth, CellHeight int32
	Cells                 [][]*Cell
}

func InitGrid(row, col int, w, h int32, seed []bool) Grid {
	cells := make([][]*Cell, col)
	for c := 0; c < col; c++ {
		cellRow := make([]*Cell, row)
		for r := 0; r < row; r++ {
			//https://en.wikipedia.org/wiki/Row-_and_column-major_order
			if seed[c*col+r] {
				cellRow[r] = NewAliveCell(r, c)
			} else {
				cellRow[r] = NewDeadCell(r, c)
			}
		}
		cells[c] = cellRow
	}
	return Grid{row, col, w, h, cells}
}

func (g *Grid) Render(renderer *Renderer) {
	for _, cellRow := range g.Cells {
		for _, c := range cellRow {
			c.Render(renderer, g.CellWidth, g.CellHeight)
		}
	}
}

func (g *Grid) GetCell(x, y int) *Cell {
	return g.Cells[y][x]
}

func (g *Grid) GetNeighbors(c *Cell) []*Cell {
	return g.getNeighbors(c.X, c.Y)
}
func (g *Grid) getNeighbors(x, y int) []*Cell {
	neighbors := make([]*Cell, 0)
	for _, dir := range AllDirections {
		dx, dy := x+dir.x, y+dir.y
		if g.isIn(dx, dy) {
			neighbors = append(neighbors, g.GetCell(dx, dy))
		}
	}
	return neighbors
}

func (g *Grid) Update() {
	for _, cellRow := range g.Cells {
		for _, cell := range cellRow {
			cell.Update(g.GetNeighbors(cell))
		}
	}
}

func (g *Grid) isIn(x, y int) bool {
	return (x >= 0 && x < g.Width) && (y >= 0 && y < g.Height)
}
