package models

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CharacterInfo struct {
	Key            int
	ImageKey       int
	DefaultTexture *sdl.Rect
	DownTextures   []*sdl.Rect
	LeftTextures   []*sdl.Rect
	RightTextures  []*sdl.Rect
	UpTextures     []*sdl.Rect
}
