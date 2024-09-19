package main

import (
	"github.com/yosoku/maze/internal/application"
	"github.com/yosoku/maze/internal/config"
)

func main() {
	cfg := config.Init(false).Get()
	app := application.Init(&cfg)
	app.Run()
}
