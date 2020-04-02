package screens

import (
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const CURSOR_SIZE = 24

type GameScreen struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	resourceManager managers.ResourceManager
	audioManager    *managers.AudioManager
	world           World
	currentCursor   int
	cursorPosition  *sdl.Point
}

func NewGameScreen(window *sdl.Window, renderer *sdl.Renderer) *GameScreen {
	resourceManager := managers.NewResourceManager()
	g := &GameScreen{
		window,
		renderer,
		resourceManager,
		managers.NewAudioManager(),
		NewWorld(window, &resourceManager, renderer, 250, 250),
		utils.NORMAL_CURSOR,
		&sdl.Point{},
	}

	// Load players textures
	g.resourceManager.LoadTexture(utils.IMG_ACTOR1, renderer)
	g.resourceManager.LoadTexture(utils.CURSORS, renderer)

	// Load audios
	g.resourceManager.LoadAudio(utils.THEME1)

	// Play music
	// g.audioManager.PlayMusic(g.resourceManager.GetAudio(utils.THEME1), 1)

	return g
}

func (g *GameScreen) HandleEvents(event sdl.Event) bool {
	running := true

	switch t := event.(type) {
	case *sdl.MouseMotionEvent:
		g.cursorPosition.X = t.X
		g.cursorPosition.Y = t.Y
	}

	g.world.HandleEvents(event)

	return running
}

func (g *GameScreen) updateMouse() {
	x := int(float64(g.cursorPosition.X)/float64(TileSize) + float64(g.world.Camera.X)/float64(TileSize))
	y := int(float64(g.cursorPosition.Y)/float64(TileSize) + float64(g.world.Camera.Y)/float64(TileSize))

	if g.world.worldMap.MapElementArray[x][y] != 0 {
		g.currentCursor = utils.HARVEST_CURSOR
	} else {
		g.currentCursor = utils.NORMAL_CURSOR
	}
}

func (g *GameScreen) Update() {
	g.world.Update()
	g.updateMouse()
}

func (g *GameScreen) renderCursor() {
	cursorInfo := utils.CursorTextureInfo[g.currentCursor]

	g.renderer.Copy(
		g.resourceManager.GetTexture(utils.CURSORS),
		cursorInfo.Src,
		&sdl.Rect{X: g.cursorPosition.X, Y: g.cursorPosition.Y, W: CURSOR_SIZE, H: CURSOR_SIZE},
	)
}

func (g *GameScreen) Render() {
	width, height := g.window.GetSize()
	g.renderer.Clear()
	g.renderer.SetDrawColor(255, 0, 0, 255)
	g.renderer.FillRect(&sdl.Rect{W: width, H: height})
	g.world.Render()
	g.renderCursor()
	g.renderer.Present()
}
