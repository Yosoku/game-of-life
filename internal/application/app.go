package application

import "C"
import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/yosoku/maze/internal/config"
	"github.com/yosoku/maze/internal/game"
	"github.com/yosoku/maze/internal/types"
	"log/slog"
	"math/rand"
	"time"
)

type Application struct {
	Window  *sdl.Window
	Running bool
	Game    *game.GameOfLife
}

func Init(cfg *config.ApplicationConfig) *Application {
	app := Application{}
	app.init(cfg)
	return &app
}

func (app *Application) init(cfg *config.ApplicationConfig) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow(cfg.WindowConfig.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		cfg.WindowConfig.Width, cfg.WindowConfig.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	app.Game = game.InitGame(cfg)
	app.Window = window
	app.Running = false
}

func (app *Application) Run() {
	app.Running = true
	defer sdl.Quit()
	defer app.Window.Destroy()

	const fps = 1
	lastTick := time.Now()
	for app.Running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			app.HandleEvent(event)
		}
		app.Render()
		now := time.Now()
		if now.Sub(lastTick) > (1/fps)*time.Second {
			app.Update()
			lastTick = now
		}
	}
	slog.Info("App quit", "application", *app)
}
func (app *Application) HandleEvent(event sdl.Event) {
	slog.Debug("Polled", "type", event.GetType(), "event", event)
	switch event.(type) {
	case *sdl.QuitEvent:
		app.Running = false
	}
}
func (app *Application) Render() {
	surface, err := app.Window.GetSurface()
	if err != nil {
		panic(err)
	}
	renderer := &types.Renderer{surface, make(map[string]any)} // inefficient as fuck
	// Fill a black screen
	surface.FillRect(nil, 0)

	app.Game.Render(renderer)

	e := app.Window.UpdateSurface()
	if e != nil {
		panic(e)
	}
}

func (app *Application) Update() {
	app.Game.Update()
}

// Debugging epilepsy
func randomRender(surface *sdl.Surface) {
	if rand.Intn(100) < 50 {
		return
	}
	rect := sdl.Rect{0, 0, 1920, 1080}
	colour := sdl.Color{200, 30, 120, 50}
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
}
