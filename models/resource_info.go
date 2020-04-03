package models

import "github.com/veandco/go-sdl2/sdl"

const (
	WOOD int = 1
)

type ResourceInfo struct {
	Key        int
	ImageKey   int
	Src        *sdl.Rect
	ResourceId int
}
