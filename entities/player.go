package entities

import (
	"math"

	"github.com/kingdoom/managers"
	"github.com/kingdoom/models"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	renderer        *sdl.Renderer
	resourceManager *managers.ResourceManager
	CharacterInfo   *models.CharacterInfo
	Pos             *sdl.Point
	posToGo         *sdl.Point
	speed           float64
}

func NewPlayer(renderer *sdl.Renderer, resourceManager *managers.ResourceManager, characterInfo *models.CharacterInfo, x int32, y int32) *Player {
	p := &Player{
		renderer,
		resourceManager,
		characterInfo,
		&sdl.Point{X: x, Y: y},
		&sdl.Point{X: x, Y: y},
		4,
	}

	return p
}

func (p *Player) OnClickToMove(mouse *sdl.MouseButtonEvent, camera *sdl.Rect) bool {
	p.posToGo.X = mouse.X + camera.X - p.CharacterInfo.Width/2
	p.posToGo.Y = mouse.Y + camera.Y - p.CharacterInfo.Height/2

	return true
}

func (p *Player) move() {
	if p.posToGo.X != p.Pos.X && p.posToGo.Y != p.Pos.Y {
		dx := float64(p.Pos.X - p.posToGo.X)
		dy := float64(p.Pos.Y - p.posToGo.Y)

		hypotenuse := math.Sqrt(dx*dx + dy*dy)
		dx /= hypotenuse
		dy /= hypotenuse
		dx *= p.speed
		dy *= p.speed

		p.Pos.X -= int32(dx)
		p.Pos.Y -= int32(dy)

		// Set the right position to avoid the problem of float to int. To be improved
		if math.Abs(dx) > math.Abs(float64(p.Pos.X-p.posToGo.X)) {
			p.Pos.X = p.posToGo.X
		}

		if math.Abs(dy) > math.Abs(float64(p.Pos.Y-p.posToGo.Y)) {
			p.Pos.Y = p.posToGo.Y
		}
	}
}

func (p *Player) Update() {
	p.move()
}

func (p *Player) Render(camera *sdl.Rect) {
	p.renderer.Copy(
		p.resourceManager.GetTexture(p.CharacterInfo.ImageKey),
		p.CharacterInfo.Src,
		&sdl.Rect{X: p.Pos.X - camera.X, Y: p.Pos.Y - camera.Y, W: p.CharacterInfo.Width, H: p.CharacterInfo.Height},
	)
}
