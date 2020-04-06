package entities

import (
	"github.com/kingdoom/managers"
	"github.com/kingdoom/models"
	"github.com/veandco/go-sdl2/sdl"
)

const BUILD_RATE = 5

type Building struct {
	resourceManager *managers.ResourceManager
	buildingInfo    *models.BuildingInfo
	Pos             *sdl.Point
	health          int
	isBuilt         bool
	frameRate       uint32
	oldTime         uint32
}

func NewBuilding(resourceManager *managers.ResourceManager, buildingInfo *models.BuildingInfo, pos *sdl.Point) *Building {
	return &Building{
		resourceManager,
		buildingInfo,
		pos,
		0,
		false,
		100,
		0,
	}
}

func (b *Building) isUserInBuilding(pos *sdl.Point, camera *sdl.Rect) bool {
	if pos.X >= b.Pos.X-camera.X && pos.X <= b.Pos.X-camera.X+b.buildingInfo.Width &&
		pos.Y >= b.Pos.Y-camera.Y && pos.Y <= b.Pos.Y-camera.Y+b.buildingInfo.Height {
		return true
	}

	return false
}

func (b *Building) construction(pos *sdl.Point, camera *sdl.Rect, isPlayerMoving bool) {
	if !b.isBuilt && b.isUserInBuilding(pos, camera) && !isPlayerMoving {
		// The player is building
		if b.oldTime+b.frameRate > sdl.GetTicks() {
			return
		}

		b.oldTime = sdl.GetTicks()

		// Add health to the building
		b.health += BUILD_RATE

		if b.health >= b.buildingInfo.Health {
			b.health = b.buildingInfo.Health
			b.isBuilt = true
		}
	}
}

func (b *Building) IsInside(pos *sdl.Point, camera *sdl.Rect, isPlayerMoving bool) bool {
	return b.isUserInBuilding(pos, camera) && !isPlayerMoving && b.isBuilt
}

func (b *Building) Update(pos *sdl.Point, camera *sdl.Rect, isPlayerMoving bool) {
	b.construction(pos, camera, isPlayerMoving)
}

func (b *Building) getCurrentTextureRect() *sdl.Rect {
	if !b.isBuilt {
		step := b.buildingInfo.Health / len(b.buildingInfo.Textures)
		index := b.health / step
		return b.buildingInfo.Textures[index]
	} else {
		return b.buildingInfo.Textures[len(b.buildingInfo.Textures)-1]
	}
}

func (b *Building) Render(renderer *sdl.Renderer, pos *sdl.Point, isPlayerMoving bool, camera *sdl.Rect) {
	if b.Pos.X+b.getCurrentTextureRect().W > camera.X && b.Pos.Y+b.getCurrentTextureRect().H > camera.Y &&
		b.Pos.X < camera.X+camera.W && b.Pos.Y < camera.Y+camera.H {
		renderer.Copy(
			b.resourceManager.GetTexture(b.buildingInfo.ImageKey),
			b.getCurrentTextureRect(),
			&sdl.Rect{
				X: b.Pos.X - camera.X,
				Y: b.Pos.Y - camera.Y,
				W: b.buildingInfo.Width,
				H: b.buildingInfo.Height,
			},
		)

		// Displays the health bar
		if !b.isBuilt && b.isUserInBuilding(pos, camera) && !isPlayerMoving {
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.DrawRect(&sdl.Rect{
				X: b.Pos.X - camera.X,
				Y: b.Pos.Y - camera.Y + b.buildingInfo.Height,
				H: HEIGHT_RECT_LIFE,
				W: b.buildingInfo.Width,
			})

			percentWidthLife := float64(b.health) / float64(b.buildingInfo.Health)

			renderer.FillRect(&sdl.Rect{
				X: b.Pos.X - camera.X,
				Y: b.Pos.Y - camera.Y + b.buildingInfo.Height,
				H: HEIGHT_RECT_LIFE,
				W: int32(float64(b.buildingInfo.Width) * percentWidthLife),
			})
		}
	}
}
