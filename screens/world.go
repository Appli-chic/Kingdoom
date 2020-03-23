package screens

import (
	"fmt"
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

const TileSize = 48

type World struct {
	resourceManager *managers.ResourceManager
	renderer        *sdl.Renderer
	mapArray        [][]int
	camera          *sdl.Rect
}

func NewWorld(window *sdl.Window, resourceManager *managers.ResourceManager, renderer *sdl.Renderer, width int, height int) World {
	matrix := make([][]int, width)
	rows := make([]int, width*height)
	for i := 0; i < width; i++ {
		matrix[i] = rows[i*height : (i+1)*height]
	}

	windowWidth, windowHeight := window.GetSize()
	w := World{resourceManager, renderer, matrix, &sdl.Rect{W: windowWidth, H: windowHeight}}
	w.initMap()

	w.resourceManager.LoadTexture(utils.OUTSIDE2, renderer)
	return w
}

func (w *World) initMap() {
	//for i := 0; i < len(w.mapArray); i++ {
	//	for j := 0; j < len(w.mapArray[0]); j++ {
	//
	//	}
	//}
}

func (w *World) HandleEvents(event sdl.Event) bool {

	switch t := event.(type) {
	case *sdl.KeyboardEvent:
		switch t.Keysym.Sym {
		case sdl.K_UP:
			w.camera.Y += 5
		case sdl.K_DOWN:
			w.camera.Y -= 5
		case sdl.K_LEFT:
			w.camera.X += 5
		case sdl.K_RIGHT:
			w.camera.X -= 5
		}
	}

	return true
}

func (w *World) Update() {

}

func (w *World) Render() {
	plainInfo := utils.TextureInfo[utils.PLAIN]

	for i := 0; i < len(w.mapArray); i++ {
		for j := 0; j < len(w.mapArray[0]); j++ {
			err := w.renderer.Copy(w.resourceManager.GetTexture(plainInfo.ImageKey), plainInfo.Src, &sdl.Rect{X: int32(TileSize * i) + w.camera.X, Y: int32(TileSize * j) + w.camera.Y, W: plainInfo.Width, H: plainInfo.Height})

			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to copy: %s\n", err)
			}
		}
	}
}
