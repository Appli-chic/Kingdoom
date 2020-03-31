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
	PLAIN int = 0

	DIRT                     int = 20
	DIRT_PLAIN_LEFT          int = 21
	DIRT_PLAIN_RIGHT         int = 22
	DIRT_PLAIN_DOWN          int = 23
	DIRT_PLAIN_UP            int = 24
	DIRT_PLAIN_LEFT_UP       int = 25
	DIRT_PLAIN_RIGHT_UP      int = 26
	DIRT_PLAIN_LEFT_BOTTOM   int = 27
	DIRT_PLAIN_RIGHT_BOTTOM  int = 28
	DIRT_CORNER_LEFT_UP      int = 29
	DIRT_CORNER_RIGHT_UP     int = 30
	DIRT_CORNER_LEFT_BOTTOM  int = 31
	DIRT_CORNER_RIGHT_BOTTOM int = 32

	WATER                     int = 40
	WATER_PLAIN_LEFT          int = 41
	WATER_PLAIN_RIGHT         int = 42
	WATER_PLAIN_DOWN          int = 43
	WATER_PLAIN_UP            int = 44
	WATER_PLAIN_LEFT_UP       int = 45
	WATER_PLAIN_RIGHT_UP      int = 46
	WATER_PLAIN_LEFT_BOTTOM   int = 47
	WATER_PLAIN_RIGHT_BOTTOM  int = 48
	WATER_CORNER_LEFT_UP      int = 49
	WATER_CORNER_RIGHT_UP     int = 50
	WATER_CORNER_LEFT_BOTTOM  int = 51
	WATER_CORNER_RIGHT_BOTTOM int = 52

	// Water animation
	WATER2                     int = 1040
	WATER2_PLAIN_LEFT          int = 1041
	WATER2_PLAIN_RIGHT         int = 1042
	WATER2_PLAIN_DOWN          int = 1043
	WATER2_PLAIN_UP            int = 1044
	WATER2_PLAIN_LEFT_UP       int = 1045
	WATER2_PLAIN_RIGHT_UP      int = 1046
	WATER2_PLAIN_LEFT_BOTTOM   int = 1047
	WATER2_PLAIN_RIGHT_BOTTOM  int = 1048
	WATER2_CORNER_LEFT_UP      int = 1049
	WATER2_CORNER_RIGHT_UP     int = 1050
	WATER2_CORNER_LEFT_BOTTOM  int = 1051
	WATER2_CORNER_RIGHT_BOTTOM int = 1052

	WATER3                     int = 2040
	WATER3_PLAIN_LEFT          int = 2041
	WATER3_PLAIN_RIGHT         int = 2042
	WATER3_PLAIN_DOWN          int = 2043
	WATER3_PLAIN_UP            int = 2044
	WATER3_PLAIN_LEFT_UP       int = 2045
	WATER3_PLAIN_RIGHT_UP      int = 2046
	WATER3_PLAIN_LEFT_BOTTOM   int = 2047
	WATER3_PLAIN_RIGHT_BOTTOM  int = 2048
	WATER3_CORNER_LEFT_UP      int = 2049
	WATER3_CORNER_RIGHT_UP     int = 2050
	WATER3_CORNER_LEFT_BOTTOM  int = 2051
	WATER3_CORNER_RIGHT_BOTTOM int = 2052
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
	PLAIN: &models.TextureInfo{Key: PLAIN, ImageKey: OUTSIDE2, Src: &sdl.Rect{W: 48, H: 48}},

	DIRT:                     &models.TextureInfo{Key: DIRT, ImageKey: OUTSIDE2, Src: &sdl.Rect{Y: 192, W: 48, H: 48}},
	DIRT_PLAIN_LEFT:          &models.TextureInfo{Key: DIRT_PLAIN_LEFT, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 96, Y: 72, W: 48, H: 48}},
	DIRT_PLAIN_RIGHT:         &models.TextureInfo{Key: DIRT_PLAIN_RIGHT, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 144, Y: 72, W: 48, H: 48}},
	DIRT_PLAIN_DOWN:          &models.TextureInfo{Key: DIRT_PLAIN_DOWN, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 120, Y: 96, W: 48, H: 48}},
	DIRT_PLAIN_UP:            &models.TextureInfo{Key: DIRT_PLAIN_UP, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 120, Y: 48, W: 48, H: 48}},
	DIRT_PLAIN_LEFT_UP:       &models.TextureInfo{Key: DIRT_PLAIN_LEFT_UP, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 96, Y: 48, W: 48, H: 48}},
	DIRT_PLAIN_RIGHT_UP:      &models.TextureInfo{Key: DIRT_PLAIN_RIGHT_UP, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 144, Y: 48, W: 48, H: 48}},
	DIRT_PLAIN_LEFT_BOTTOM:   &models.TextureInfo{Key: DIRT_PLAIN_LEFT_BOTTOM, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 96, Y: 96, W: 48, H: 48}},
	DIRT_PLAIN_RIGHT_BOTTOM:  &models.TextureInfo{Key: DIRT_PLAIN_RIGHT_BOTTOM, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 144, Y: 96, W: 48, H: 48}},
	DIRT_CORNER_LEFT_UP:      &models.TextureInfo{Key: DIRT_CORNER_LEFT_UP, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 768, W: 48, H: 48}},
	DIRT_CORNER_RIGHT_UP:     &models.TextureInfo{Key: DIRT_CORNER_RIGHT_UP, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 816, W: 48, H: 48}},
	DIRT_CORNER_LEFT_BOTTOM:  &models.TextureInfo{Key: DIRT_CORNER_LEFT_BOTTOM, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 768, Y: 48, W: 48, H: 48}},
	DIRT_CORNER_RIGHT_BOTTOM: &models.TextureInfo{Key: DIRT_CORNER_RIGHT_BOTTOM, ImageKey: OUTSIDE2, Src: &sdl.Rect{X: 816, Y: 48, W: 48, H: 48}},

	WATER:                     &models.TextureInfo{Key: WATER, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 24, Y: 72, W: 48, H: 48}},
	WATER_PLAIN_LEFT:          &models.TextureInfo{Key: WATER_PLAIN_LEFT, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 72, W: 48, H: 48}},
	WATER_PLAIN_RIGHT:         &models.TextureInfo{Key: WATER_PLAIN_RIGHT, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 48, Y: 72, W: 48, H: 48}},
	WATER_PLAIN_DOWN:          &models.TextureInfo{Key: WATER_PLAIN_DOWN, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 24, Y: 96, W: 48, H: 48}},
	WATER_PLAIN_UP:            &models.TextureInfo{Key: WATER_PLAIN_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 24, Y: 48, W: 48, H: 48}},
	WATER_PLAIN_LEFT_UP:       &models.TextureInfo{Key: WATER_PLAIN_LEFT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 48, W: 48, H: 48}},
	WATER_PLAIN_RIGHT_UP:      &models.TextureInfo{Key: WATER_PLAIN_RIGHT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 48, Y: 48, W: 48, H: 48}},
	WATER_PLAIN_LEFT_BOTTOM:   &models.TextureInfo{Key: WATER_PLAIN_LEFT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 96, W: 48, H: 48}},
	WATER_PLAIN_RIGHT_BOTTOM:  &models.TextureInfo{Key: WATER_PLAIN_RIGHT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 48, Y: 96, W: 48, H: 48}},
	WATER_CORNER_LEFT_UP:      &models.TextureInfo{Key: WATER_CORNER_LEFT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 768, W: 48, H: 48}},
	WATER_CORNER_RIGHT_UP:     &models.TextureInfo{Key: WATER_CORNER_RIGHT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 816, W: 48, H: 48}},
	WATER_CORNER_LEFT_BOTTOM:  &models.TextureInfo{Key: WATER_CORNER_LEFT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 768, Y: 48, W: 48, H: 48}},
	WATER_CORNER_RIGHT_BOTTOM: &models.TextureInfo{Key: WATER_CORNER_RIGHT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 816, Y: 48, W: 48, H: 48}},

	// Water animation
	WATER2:                     &models.TextureInfo{Key: WATER2, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 120, Y: 72, W: 48, H: 48}},
	WATER2_PLAIN_LEFT:          &models.TextureInfo{Key: WATER2_PLAIN_LEFT, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 72, W: 48, H: 48}},
	WATER2_PLAIN_RIGHT:         &models.TextureInfo{Key: WATER2_PLAIN_RIGHT, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 144, Y: 72, W: 48, H: 48}},
	WATER2_PLAIN_DOWN:          &models.TextureInfo{Key: WATER2_PLAIN_DOWN, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 120, Y: 96, W: 48, H: 48}},
	WATER2_PLAIN_UP:            &models.TextureInfo{Key: WATER2_PLAIN_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 120, Y: 48, W: 48, H: 48}},
	WATER2_PLAIN_LEFT_UP:       &models.TextureInfo{Key: WATER2_PLAIN_LEFT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 48, W: 48, H: 48}},
	WATER2_PLAIN_RIGHT_UP:      &models.TextureInfo{Key: WATER2_PLAIN_RIGHT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 144, Y: 48, W: 48, H: 48}},
	WATER2_PLAIN_LEFT_BOTTOM:   &models.TextureInfo{Key: WATER2_PLAIN_LEFT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 96, W: 48, H: 48}},
	WATER2_PLAIN_RIGHT_BOTTOM:  &models.TextureInfo{Key: WATER2_PLAIN_RIGHT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 144, Y: 96, W: 48, H: 48}},
	WATER2_CORNER_LEFT_UP:      &models.TextureInfo{Key: WATER2_CORNER_LEFT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 864, W: 48, H: 48}},
	WATER2_CORNER_RIGHT_UP:     &models.TextureInfo{Key: WATER2_CORNER_RIGHT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 912, W: 48, H: 48}},
	WATER2_CORNER_LEFT_BOTTOM:  &models.TextureInfo{Key: WATER2_CORNER_LEFT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 864, Y: 48, W: 48, H: 48}},
	WATER2_CORNER_RIGHT_BOTTOM: &models.TextureInfo{Key: WATER2_CORNER_RIGHT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 912, Y: 48, W: 48, H: 48}},

	WATER3:                     &models.TextureInfo{Key: WATER3, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 120, Y: 72, W: 48, H: 48}},
	WATER3_PLAIN_LEFT:          &models.TextureInfo{Key: WATER3_PLAIN_LEFT, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 72, W: 48, H: 48}},
	WATER3_PLAIN_RIGHT:         &models.TextureInfo{Key: WATER3_PLAIN_RIGHT, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 240, Y: 72, W: 48, H: 48}},
	WATER3_PLAIN_DOWN:          &models.TextureInfo{Key: WATER3_PLAIN_DOWN, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 216, Y: 96, W: 48, H: 48}},
	WATER3_PLAIN_UP:            &models.TextureInfo{Key: WATER3_PLAIN_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 216, Y: 48, W: 48, H: 48}},
	WATER3_PLAIN_LEFT_UP:       &models.TextureInfo{Key: WATER3_PLAIN_LEFT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 48, W: 48, H: 48}},
	WATER3_PLAIN_RIGHT_UP:      &models.TextureInfo{Key: WATER3_PLAIN_RIGHT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 240, Y: 48, W: 48, H: 48}},
	WATER3_PLAIN_LEFT_BOTTOM:   &models.TextureInfo{Key: WATER3_PLAIN_LEFT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{Y: 96, W: 48, H: 48}},
	WATER3_PLAIN_RIGHT_BOTTOM:  &models.TextureInfo{Key: WATER3_PLAIN_RIGHT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 240, Y: 96, W: 48, H: 48}},
	WATER3_CORNER_LEFT_UP:      &models.TextureInfo{Key: WATER3_CORNER_LEFT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 960, W: 48, H: 48}},
	WATER3_CORNER_RIGHT_UP:     &models.TextureInfo{Key: WATER3_CORNER_RIGHT_UP, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 1008, W: 48, H: 48}},
	WATER3_CORNER_LEFT_BOTTOM:  &models.TextureInfo{Key: WATER3_CORNER_LEFT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 960, Y: 48, W: 48, H: 48}},
	WATER3_CORNER_RIGHT_BOTTOM: &models.TextureInfo{Key: WATER3_CORNER_RIGHT_BOTTOM, ImageKey: OUTSIDE1, Src: &sdl.Rect{X: 1008, Y: 48, W: 48, H: 48}},
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
