package screens

import (
	"github.com/kingdoom/entities"
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const TileSize = 48

type World struct {
	resourceManager *managers.ResourceManager
	renderer        *sdl.Renderer
	worldMap        *Map
	camera          *sdl.Rect
	player          *entities.Character
}

func NewWorld(window *sdl.Window, resourceManager *managers.ResourceManager, renderer *sdl.Renderer, width int, height int) World {
	windowWidth, windowHeight := window.GetSize()
	w := World{
		resourceManager,
		renderer,
		NewMap(width, height),
		&sdl.Rect{W: windowWidth, H: windowHeight},
		entities.NewPlayer(renderer, resourceManager, utils.CharacterTextureInfo[utils.ACTOR1], 50, 50),
	}

	// Loading textures
	w.resourceManager.LoadTexture(utils.OUTSIDE1, renderer)
	w.resourceManager.LoadTexture(utils.OUTSIDE2, renderer)

	return w
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
	w.worldMap.Update()
}

func (w *World) centerCamera() {
	w.camera.X = int32(w.player.Pos.X) + w.player.GetWidth()/2 - w.camera.W/2
	w.camera.Y = int32(w.player.Pos.Y) + w.player.GetHeight()/2 - w.camera.H/2

	// Manage map corners
	if w.camera.X < 1*TileSize {
		w.camera.X = 1 * TileSize
	}

	if w.camera.Y < 1*TileSize {
		w.camera.Y = 1 * TileSize
	}

	if w.camera.X+w.camera.W > int32(len(w.worldMap.MapArray)*TileSize)-1*TileSize {
		w.camera.X = int32(len(w.worldMap.MapArray)*TileSize) - w.camera.W - 1*TileSize
	}

	if w.camera.Y+w.camera.H > int32(len(w.worldMap.MapArray[0])*TileSize)-1*TileSize {
		w.camera.Y = int32(len(w.worldMap.MapArray[0])*TileSize) - w.camera.H - 1*TileSize
	}
}

func (w *World) Render() {
	w.worldMap.Render(w.camera, w.resourceManager, w.renderer)
	w.player.Render(w.camera)
}
