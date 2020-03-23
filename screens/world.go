package screens

import (
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const TILE_SIZE = 48

type World struct {
	resourceManager *managers.ResourceManager
	renderer        *sdl.Renderer
	mapArray        [][]int
}

func NewWorld(resourceManager *managers.ResourceManager, renderer *sdl.Renderer, width int, height int) World {
	matrix := make([][]int, width)
	rows := make([]int, width*height)
	for i := 0; i < width; i++ {
		matrix[i] = rows[i*height : (i+1)*height]
	}

	w := World{resourceManager, renderer, matrix}
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

func (w *World) HandleEvents() bool {
	return true
}

func (w *World) Update() {

}

func (w *World) Render() {
	plainInfo := utils.TextureInfo[utils.PLAIN]

	for i := 0; i < len(w.mapArray); i++ {
		for j := 0; j < len(w.mapArray[0]); j++ {
			w.renderer.Copy(w.resourceManager.GetTexture(plainInfo.ImageKey), plainInfo.Src, &sdl.Rect{X: int32(TILE_SIZE * i), Y: int32(TILE_SIZE * j), W: plainInfo.Width, H: plainInfo.Height})
		}
	}
}
