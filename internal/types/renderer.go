package types

import "github.com/veandco/go-sdl2/sdl"

// Maybe down the line i'll figure out a way to not use the surface to render
// So for now just wrap this up in a struct
type Renderer struct {
	Surface  *sdl.Surface
	Metadata map[string]any
}
