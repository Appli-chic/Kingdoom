package models

import "github.com/veandco/go-sdl2/sdl"

type TextureInfo struct {
	Key      int
	ImageKey int
	Src      *sdl.Rect
}
