package screens

import (
	"fmt"
	"os"

	"github.com/kingdoom/entities"
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const TileSize = 48

type World struct {
	resourceManager *managers.ResourceManager
	renderer        *sdl.Renderer
	mapArray        [][]int
	camera          *sdl.Rect
	player          *entities.Player
}

func NewWorld(window *sdl.Window, resourceManager *managers.ResourceManager, renderer *sdl.Renderer, width int, height int) World {
	matrix := make([][]int, width)
	rows := make([]int, width*height)
	for i := 0; i < width; i++ {
		matrix[i] = rows[i*height : (i+1)*height]
	}

	windowWidth, windowHeight := window.GetSize()
	w := World{
		resourceManager,
		renderer,
		matrix,
		&sdl.Rect{W: windowWidth, H: windowHeight},
		entities.NewPlayer(renderer, resourceManager, utils.CharacterTextureInfo[utils.ACTOR1], 50, 50),
	}
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
	case *sdl.MouseButtonEvent:
		if t.State != sdl.PRESSED {
			if t.Button == sdl.BUTTON_LEFT {
				w.player.OnClickToMove(t, w.camera)
			}
		}
	}

	return true
}

func (w *World) Update() {
	w.player.Update()
	w.centerCamera()
}

func (w *World) centerCamera() {
	w.camera.X = int32(w.player.Pos.X) + w.player.CharacterInfo.Width/2 - w.camera.W/2
	w.camera.Y = int32(w.player.Pos.Y) + w.player.CharacterInfo.Height/2 - w.camera.H/2
}

func (w *World) Render() {
	plainInfo := utils.GroundTextureInfo[utils.PLAIN]

	for i := 0; i < len(w.mapArray); i++ {
		for j := 0; j < len(w.mapArray[0]); j++ {
			err := w.renderer.Copy(
				w.resourceManager.GetTexture(plainInfo.ImageKey),
				plainInfo.Src,
				&sdl.Rect{
					X: int32(TileSize*i) - w.camera.X,
					Y: int32(TileSize*j) - w.camera.Y,
					W: plainInfo.Width,
					H: plainInfo.Height,
				},
			)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to copy: %s\n", err)
			}
		}
	}

	w.player.Render(w.camera)
}
