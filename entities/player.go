package entities

import (
	"github.com/kingdoom/managers"
	"github.com/kingdoom/models"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Player struct {
	renderer        *sdl.Renderer
	resourceManager *managers.ResourceManager
	CharacterInfo   *models.CharacterInfo
	Pos             *sdl.Point
	posToGo         *sdl.Point
	posStart        *sdl.Point
	speed           float64
}

func NewPlayer(renderer *sdl.Renderer, resourceManager *managers.ResourceManager, characterInfo *models.CharacterInfo, x int32, y int32) *Player {
	p := &Player{
		renderer,
		resourceManager,
		characterInfo,
		&sdl.Point{X: x, Y: y},
		&sdl.Point{X: x, Y: y},
		&sdl.Point{X: x, Y: y},
		4,
	}

	return p
}

func (p *Player) OnClickToMove(mouse *sdl.MouseButtonEvent, camera *sdl.Rect) bool {
	p.posToGo.X = mouse.X + camera.X - p.CharacterInfo.Width/2
	p.posToGo.Y = mouse.Y + camera.Y - p.CharacterInfo.Height/2

	p.posStart.X = p.Pos.X
	p.posStart.Y = p.Pos.Y

	return true
}

func (p *Player) move() {
	if p.posToGo.X != p.Pos.X && p.posToGo.Y != p.Pos.Y {
		speed := p.speed * 0.1
		hypotenuse := math.Sqrt(math.Abs(float64(p.posStart.X-p.posToGo.X)) + math.Abs(float64(p.posStart.Y-p.posToGo.Y)))

		dx := math.Abs(float64(p.posStart.X-p.posToGo.X)) / hypotenuse * speed
		dy := math.Abs(float64(p.posStart.Y-p.posToGo.Y)) / hypotenuse * speed

		if p.posToGo.X > p.Pos.X {
			if p.Pos.X+int32(dx) > p.posToGo.X {
				p.Pos.X = p.posToGo.X
			} else {
				p.Pos.X += int32(dx)
			}
		} else {
			if p.Pos.X-int32(dx) < p.posToGo.X {
				p.Pos.X = p.posToGo.X
			} else {
				p.Pos.X -= int32(dx)
			}
		}

		if p.posToGo.Y > p.Pos.Y {
			if p.Pos.Y+int32(dy) > p.posToGo.Y {
				p.Pos.Y = p.posToGo.Y
			} else {
				p.Pos.Y += int32(dy)
			}
		} else {
			if p.Pos.Y-int32(dy) < p.posToGo.Y {
				p.Pos.Y = p.posToGo.Y
			} else {
				p.Pos.Y -= int32(dy)
			}
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
