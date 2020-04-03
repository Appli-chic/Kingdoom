package screens

import (
	"github.com/kingdoom/entities"
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type World struct {
	resourceManager *managers.ResourceManager
	renderer        *sdl.Renderer
	worldMap        *Map
	Camera          *sdl.Rect
	player          *entities.Character
}

func NewWorld(window *sdl.Window, resourceManager *managers.ResourceManager, renderer *sdl.Renderer, width int, height int) World {
	windowWidth, windowHeight := window.GetSize()
	w := World{
		resourceManager,
		renderer,
		NewMap(width, height),
		&sdl.Rect{W: windowWidth, H: windowHeight},
		entities.NewPlayer(renderer, resourceManager, utils.CharacterTextureInfo[utils.ACTOR1], 50, 50, true),
	}

	// Loading textures
	w.resourceManager.LoadTexture(utils.OUTSIDE1, renderer)
	w.resourceManager.LoadTexture(utils.OUTSIDE2, renderer)
	w.resourceManager.LoadTexture(utils.OUTSIDEB, renderer)

	return w
}

func (w *World) HandleEvents(event sdl.Event) bool {
	switch t := event.(type) {
	case *sdl.MouseButtonEvent:
		if t.State != sdl.PRESSED {
			if t.Button == sdl.BUTTON_LEFT {
				w.player.OnClickToMove(t, w.Camera)
			}
		}
	}

	return true
}

func (w *World) Update() {
	w.player.Update(w.worldMap.MapResourceArray)
	w.centerCamera()
	w.worldMap.Update()
}

func (w *World) centerCamera() {
	w.Camera.X = int32(w.player.Pos.X) + w.player.GetWidth()/2 - w.Camera.W/2
	w.Camera.Y = int32(w.player.Pos.Y) + w.player.GetHeight()/2 - w.Camera.H/2

	// Manage map corners
	if w.Camera.X < utils.TileSize {
		w.Camera.X = utils.TileSize
	}

	if w.Camera.Y < utils.TileSize {
		w.Camera.Y = utils.TileSize
	}

	if w.Camera.X+w.Camera.W > int32(len(w.worldMap.MapArray)*utils.TileSize)-utils.TileSize {
		w.Camera.X = int32(len(w.worldMap.MapArray)*utils.TileSize) - w.Camera.W - utils.TileSize
	}

	if w.Camera.Y+w.Camera.H > int32(len(w.worldMap.MapArray[0])*utils.TileSize)-utils.TileSize {
		w.Camera.Y = int32(len(w.worldMap.MapArray[0])*utils.TileSize) - w.Camera.H - utils.TileSize
	}
}

func (w *World) Render() {
	w.worldMap.Render(w.Camera, w.resourceManager, w.renderer)
	w.player.Render(w.Camera)
}
