package util

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

func RandomColor() sdl.Color {
	return sdl.Color{R: Random8Bit(), G: Random8Bit(), B: Random8Bit(), A: Random8Bit()}
}

func Random8Bit() uint8 {
	r := rand.Int31n(255)
	return uint8(r)
}
