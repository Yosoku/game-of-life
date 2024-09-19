package config

import (
	"log/slog"
	"os"
	"strconv"
)

type WindowConfig struct {
	Width, Height int32
	Title         string
}
type GameConfig struct {
	RandomSeed bool
	GridX      int
	GridY      int
}
type ApplicationConfig struct {
	WindowConfig WindowConfig
	GameConfig   GameConfig
}
type Config interface {
	Get() ApplicationConfig
}
type ReleaseConfig struct {
	Config
	AppConfig ApplicationConfig
}
type DevConfig struct {
	Config
	AppConfig ApplicationConfig
}

func (dc DevConfig) Get() ApplicationConfig {
	return dc.AppConfig
}

// Pretend I'm gonna release this
func (rc ReleaseConfig) Get() ApplicationConfig {
	return rc.AppConfig
}

func Init(release bool) Config {
	if release {
		return initRelease()
	}
	slog.SetLogLoggerLevel(slog.LevelDebug)
	return initDev()
}

func initDev() Config {
	gameCfg := GameConfig{
		RandomSeed: true,
		GridX:      160,
		GridY:      90,
	}
	windowCfg := WindowConfig{Width: 1920, Height: 1080, Title: "Game of life"}
	return DevConfig{AppConfig: ApplicationConfig{
		WindowConfig: windowCfg,
		GameConfig:   gameCfg,
	}}
}

func initRelease() Config {
	seed := loadEnv[bool]("G_SEED", "true", strconv.ParseBool)
	gX := loadEnv("G_GRID_X", "100", strconv.Atoi)
	gY := loadEnv("G_GRID_Y", "100", strconv.Atoi)
	wWidth := loadEnv("W_WIDTH", "1920", strconv.Atoi)
	wHeight := loadEnv("W_HEIGHT", "1080", strconv.Atoi)
	wTitle := loadEnv("W_TITLE", "AAAAAAA", func(s string) (string, error) { return s, nil })
	gameCfg := GameConfig{
		RandomSeed: seed,
		GridX:      gX,
		GridY:      gY,
	}
	windowCfg := WindowConfig{Width: int32(wWidth), Height: int32(wHeight), Title: wTitle}
	slog.Info("Initialized release config", "window", windowCfg, "game", gameCfg)
	return ReleaseConfig{AppConfig: ApplicationConfig{
		WindowConfig: windowCfg,
		GameConfig:   gameCfg,
	}}
}

type Transform[T comparable] func(str string) (T, error)

func loadEnv[T comparable](k, def string, f Transform[T]) T {
	a := load(k, def)
	out, err := f(a)
	if err != nil {
		panic(err)
	}
	return out
}

func load(key, defaultVal string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultVal
}
