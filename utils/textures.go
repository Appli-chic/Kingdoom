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
	SAND int = 20
)

var ImagesPath = map[int]string{
	OUTSIDE2: "assets/tileset/outside2.png",
}

var TextureInfo = map[int]models.TextureInfo{
	PLAIN: models.TextureInfo{Key: PLAIN, ImageKey: OUTSIDE2, Src: &sdl.Rect{W: 48, H: 48}, Width: 48, Height: 48},
	SAND: models.TextureInfo{Key: SAND, ImageKey: OUTSIDE2, Src: &sdl.Rect{Y: 192, W: 48, H: 48}, Width: 48, Height: 48},
}
