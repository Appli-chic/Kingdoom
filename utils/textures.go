package utils

import (
	"github.com/kingdoom/models"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	OUTSIDE1 int  = 0
	OUTSIDE2 int  = 1
)

const (
	PLAIN int  = 0
	DESERT int = 20
)

var TextureInfo = map[int]models.Texture{
	PLAIN: models.Texture{Key: PLAIN, ImageKey: OUTSIDE2, Src: &sdl.Rect{W: 48, H: 48}, Width: 48, Height: 48},
}
