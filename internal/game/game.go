package game

import (
	"github.com/yosoku/maze/internal/config"
	"github.com/yosoku/maze/internal/types"
	"log/slog"
	"math/rand"
)

type GameOfLife struct {
	Grid       types.Grid
	Generation int
}

func InitGame(cfg *config.ApplicationConfig) *GameOfLife {
	xs := cfg.GameConfig.GridX
	ys := cfg.GameConfig.GridY
	CellWidth := cfg.WindowConfig.Width / int32(xs)
	CellHeight := cfg.WindowConfig.Width / int32(ys)
	slog.Info("Initializing grid", "xSquares", xs, "ySquares", ys, "cellWidth", CellWidth, "cellHeight", CellHeight)
	seed := make([]bool, 0)
	if cfg.GameConfig.RandomSeed {
		seed = createSeed(xs * ys)
	}
	grid := types.InitGrid(xs, ys, CellWidth, CellHeight, seed)
	return &GameOfLife{grid, 0}
}

func createSeed(size int) []bool {
	bits := make([]bool, size)
	for i := 0; i < size; i++ {
		bits[i] = rand.Intn(100) > 90
	}
	return bits
}

func (m *GameOfLife) Render(renderer *types.Renderer) {
	m.Grid.Render(renderer)
}

func (m *GameOfLife) Update() {
	m.Grid.Update()
	m.Generation++
}
