package utils

import (
	"github.com/kingdoom/models"
	"github.com/veandco/go-sdl2/sdl"
)

// Image keys
const (
	// Ground
	OUTSIDE1 int = 0
	OUTSIDE2 int = 1

	// Characters
	IMG_ACTOR1 int = 1000
)

// Ground texture keys
const (
	PLAIN       int = 0
	DIRT        int = 20
	SAND        int = 40
	WATER_GRASS int = 60
)

// Character texture keys
const (
	ACTOR1 int = 1000
)

var ImagesPath = map[int]string{
	// Ground
	OUTSIDE1: "assets/tileset/outside1.png",
	OUTSIDE2: "assets/tileset/outside2.png",

	// Characters
	IMG_ACTOR1: "assets/character/actor1.png",
}

var GroundTextureInfo = map[int]*models.TextureInfo{
	PLAIN:       &models.TextureInfo{Key: PLAIN, ImageKey: OUTSIDE2, Src: &sdl.Rect{W: 48, H: 48}},
	DIRT:        &models.TextureInfo{Key: DIRT, ImageKey: OUTSIDE2, Src: &sdl.Rect{Y: 192, W: 48, H: 48}},
	SAND:        &models.TextureInfo{Key: SAND, ImageKey: OUTSIDE2, Src: &sdl.Rect{Y: 192, W: 48, H: 48}},
	WATER_GRASS: &models.TextureInfo{Key: WATER_GRASS, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 24, Y: 72, W: 48, H: 48}},
}

var CharacterTextureInfo = map[int]*models.CharacterInfo{
	ACTOR1: &models.CharacterInfo{
		Key:            ACTOR1,
		ImageKey:       IMG_ACTOR1,
		DefaultTexture: &sdl.Rect{X: 48, W: 48, H: 48},
		DownTextures: []*sdl.Rect{
			&sdl.Rect{W: 48, H: 48},
			&sdl.Rect{X: 48, W: 48, H: 48},
			&sdl.Rect{X: 96, W: 48, H: 48},
		},
		LeftTextures: []*sdl.Rect{
			&sdl.Rect{Y: 48, W: 48, H: 48},
			&sdl.Rect{X: 48, Y: 48, W: 48, H: 48},
			&sdl.Rect{X: 96, Y: 48, W: 48, H: 48},
		},
		RightTextures: []*sdl.Rect{
			&sdl.Rect{Y: 96, W: 48, H: 48},
			&sdl.Rect{X: 48, Y: 96, W: 48, H: 48},
			&sdl.Rect{X: 96, Y: 96, W: 48, H: 48},
		},
		UpTextures: []*sdl.Rect{
			&sdl.Rect{Y: 144, W: 48, H: 48},
			&sdl.Rect{X: 48, Y: 144, W: 48, H: 48},
			&sdl.Rect{X: 96, Y: 144, W: 48, H: 48},
		},
	},
}
