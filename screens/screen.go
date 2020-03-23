package screens

import "github.com/veandco/go-sdl2/sdl"

type Screen interface {
	HandleEvents(event sdl.Event) bool
	Update()
	Render()
}
