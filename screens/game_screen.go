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
	isBuilding      bool
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
		false,
	}

	// Loading textures
	g.resourceManager.LoadTexture(utils.OUTSIDE1, renderer)
	g.resourceManager.LoadTexture(utils.OUTSIDE2, renderer)
	g.resourceManager.LoadTexture(utils.OUTSIDEB, renderer)

	// Load players textures
	g.resourceManager.LoadTexture(utils.IMG_ACTOR1, renderer)

	// Load building textures
	g.resourceManager.LoadTexture(utils.SAWMILL, renderer)

	// Load cursor textures
	g.resourceManager.LoadTexture(utils.CURSORS, renderer)

	// Load fonts
	g.resourceManager.LoadFont(utils.FONT_FIRACODE, 18)

	// Load audios
	g.resourceManager.LoadAudio(utils.THEME1)

	// Play music
	// g.audioManager.PlayMusic(g.resourceManager.GetAudio(utils.THEME1), 1)

	return g
}

func (g *GameScreen) HandleEvents(event sdl.Event) bool {
	running := true
	isInterfaceClick := false

	switch t := event.(type) {
	case *sdl.MouseMotionEvent:
		g.cursorPosition.X = t.X
		g.cursorPosition.Y = t.Y
	case *sdl.MouseButtonEvent:
		if t.State != sdl.PRESSED {
			if t.Button == sdl.BUTTON_LEFT && g.isBuilding {
				buildingInfo := utils.BuildingTextureInfo[utils.SAWMILL]
				pos := &sdl.Point{
					X: t.X + g.world.Camera.X,
					Y: t.Y + g.world.Camera.Y,
				}

				g.world.player.Build(buildingInfo, pos)
				g.isBuilding = false
				isInterfaceClick = true
			}
		}
	case *sdl.KeyboardEvent:
		if t.Keysym.Sym == sdl.K_b {
			g.isBuilding = true
		}
	}

	g.world.HandleEvents(event, isInterfaceClick)

	return running
}

func (g *GameScreen) updateMouse() {
	x := int(float64(g.cursorPosition.X)/float64(utils.TileSize) + float64(g.world.Camera.X)/float64(utils.TileSize))
	y := int(float64(g.cursorPosition.Y)/float64(utils.TileSize) + float64(g.world.Camera.Y)/float64(utils.TileSize))

	if g.isBuilding {
		g.currentCursor = utils.BUILDING_OK_CURSOR
	} else if g.world.worldMap.MapResourceArray[x][y] != 0 {
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

func (g *GameScreen) renderBuildingSpace() {
	if g.isBuilding {
		buildingInfo := utils.BuildingTextureInfo[utils.SAWMILL]

		g.renderer.Copy(
			g.resourceManager.GetTexture(utils.SAWMILL),
			buildingInfo.Textures[0],
			&sdl.Rect{X: g.cursorPosition.X - buildingInfo.Width/2, Y: g.cursorPosition.Y - buildingInfo.Height/2, W: buildingInfo.Width, H: buildingInfo.Height},
		)
	}
}

func (g *GameScreen) Render() {
	width, height := g.window.GetSize()
	g.renderer.Clear()
	g.renderer.SetDrawColor(255, 0, 0, 255)
	g.renderer.FillRect(&sdl.Rect{W: width, H: height})

	g.world.Render()
	g.renderBuildingSpace()
	g.renderCursor()

	g.renderer.Present()
}
