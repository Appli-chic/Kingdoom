package entities

import (
	"math"

	"github.com/kingdoom/managers"
	"github.com/kingdoom/models"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	DIRECTION_DEFAULT = 0
	DIRECTION_DOWN    = 1
	DIRECTION_LEFT    = 2
	DIRECTION_RIGHT   = 3
	DIRECTION_UP      = 4
)

type Character struct {
	renderer        *sdl.Renderer
	resourceManager *managers.ResourceManager
	CharacterInfo   *models.CharacterInfo
	Pos             *sdl.Point
	posToGo         *sdl.Point
	speed           float64
	direction       int
	currentFrame    int
	frameRate       uint32
	oldTime         uint32
}

func NewPlayer(renderer *sdl.Renderer, resourceManager *managers.ResourceManager, characterInfo *models.CharacterInfo, x int32, y int32) *Character {
	c := &Character{
		renderer,
		resourceManager,
		characterInfo,
		&sdl.Point{X: x, Y: y},
		&sdl.Point{X: x, Y: y},
		50,
		DIRECTION_DEFAULT,
		0,
		100,
		0,
	}

	return c
}

func (c *Character) GetHeight() int32 {
	return c.CharacterInfo.DefaultTexture.H
}

func (c *Character) GetWidth() int32 {
	return c.CharacterInfo.DefaultTexture.W
}

func (c *Character) OnClickToMove(mouse *sdl.MouseButtonEvent, camera *sdl.Rect) bool {
	c.posToGo.X = mouse.X + camera.X - c.CharacterInfo.DefaultTexture.W/2
	c.posToGo.Y = mouse.Y + camera.Y - c.CharacterInfo.DefaultTexture.H/2

	return true
}

func (c *Character) setDirection(dx float64, dy float64) {
	isLookingVertically := math.Abs(dy) > math.Abs(dx)

	if isLookingVertically {
		// Is looking Up or Down
		if dy >= 0 {
			c.direction = DIRECTION_UP
		} else {
			c.direction = DIRECTION_DOWN
		}
	} else {
		// Is looking Left or Right
		if dx >= 0 {
			c.direction = DIRECTION_LEFT
		} else {
			c.direction = DIRECTION_RIGHT
		}
	}
}

func (c *Character) getCurrentTextureRect() *sdl.Rect {
	switch c.direction {
	case DIRECTION_DEFAULT:
		return c.CharacterInfo.DefaultTexture
	case DIRECTION_DOWN:
		return c.CharacterInfo.DownTextures[c.currentFrame]
	case DIRECTION_LEFT:
		return c.CharacterInfo.LeftTextures[c.currentFrame]
	case DIRECTION_RIGHT:
		return c.CharacterInfo.RightTextures[c.currentFrame]
	case DIRECTION_UP:
		return c.CharacterInfo.UpTextures[c.currentFrame]
	default:
		return c.CharacterInfo.DefaultTexture
	}
}

func (c *Character) animate() {
	if c.oldTime+c.frameRate > sdl.GetTicks() {
		return
	}

	c.oldTime = sdl.GetTicks()
	c.currentFrame++

	switch c.direction {
	case DIRECTION_DOWN:
		if c.currentFrame >= len(c.CharacterInfo.DownTextures) {
			c.currentFrame = 0
		}
	case DIRECTION_LEFT:
		if c.currentFrame >= len(c.CharacterInfo.LeftTextures) {
			c.currentFrame = 0
		}
	case DIRECTION_RIGHT:
		if c.currentFrame >= len(c.CharacterInfo.RightTextures) {
			c.currentFrame = 0
		}
	case DIRECTION_UP:
		if c.currentFrame >= len(c.CharacterInfo.UpTextures) {
			c.currentFrame = 0
		}
	default:
		c.currentFrame = 1
	}
}

func (c *Character) move() {
	if c.posToGo.X != c.Pos.X && c.posToGo.Y != c.Pos.Y {
		dx := float64(c.Pos.X - c.posToGo.X)
		dy := float64(c.Pos.Y - c.posToGo.Y)

		hypotenuse := math.Sqrt(dx*dx + dy*dy)
		dx /= hypotenuse
		dy /= hypotenuse
		dx *= c.speed
		dy *= c.speed

		c.setDirection(dx, dy)
		c.animate()

		c.Pos.X -= int32(dx)
		c.Pos.Y -= int32(dy)

		// Set the right position to avoid the problem of float to int. To be improved
		if math.Abs(dx) > math.Abs(float64(c.Pos.X-c.posToGo.X)) {
			c.Pos.X = c.posToGo.X
		}

		if math.Abs(dy) > math.Abs(float64(c.Pos.Y-c.posToGo.Y)) {
			c.Pos.Y = c.posToGo.Y
		}
	} else {
		// Is not moving anymore
		c.direction = DIRECTION_DEFAULT
	}
}

func (c *Character) Update() {
	c.move()
}

func (c *Character) Render(camera *sdl.Rect) {
	c.renderer.Copy(
		c.resourceManager.GetTexture(c.CharacterInfo.ImageKey),
		c.getCurrentTextureRect(),
		&sdl.Rect{X: c.Pos.X - camera.X, Y: c.Pos.Y - camera.Y, W: c.CharacterInfo.DefaultTexture.W, H: c.CharacterInfo.DefaultTexture.H},
	)
}
