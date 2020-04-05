package models

import "github.com/veandco/go-sdl2/sdl"

type BuildingInfo struct {
	ImageKey int
	Width    int32
	Height   int32
	Health   int
	Textures []*sdl.Rect
}
