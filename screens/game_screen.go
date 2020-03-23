package screens

import (
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type GameScreen struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	resourceManager managers.ResourceManager
}

func NewGameScreen(window *sdl.Window, renderer *sdl.Renderer) *GameScreen {
	g := &GameScreen{window, renderer, managers.NewResourceManager()}

	g.resourceManager.LoadImage(utils.OUTSIDE2)
	g.resourceManager.LoadTextureFromImage(utils.PLAIN, g.resourceManager.GetImage(utils.OUTSIDE2), renderer)

	return g
}

func (g *GameScreen) HandleEvents() bool {
	running := true

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			running = false
			break
		}
	}

	return running
}

func (g *GameScreen) Update() {

}

func (g *GameScreen) Render() {
	width, height := g.window.GetSize()
	textureInfo := utils.TextureInfo[utils.PLAIN]

	g.renderer.Clear()
	g.renderer.SetDrawColor(255, 0, 0, 255)
	g.renderer.FillRect(&sdl.Rect{W: width, H: height})
	g.renderer.Copy(g.resourceManager.GetTexture(utils.PLAIN), textureInfo.Src, &sdl.Rect{W: textureInfo.Width, H: textureInfo.Height})
	g.renderer.Present()
}
