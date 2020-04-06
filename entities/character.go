package entities

import (
	"math"

	"github.com/kingdoom/managers"
	"github.com/kingdoom/models"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const HEIGHT_RECT_LIFE = 10

const (
	DIRECTION_DEFAULT = 0
	DIRECTION_DOWN    = 1
	DIRECTION_LEFT    = 2
	DIRECTION_RIGHT   = 3
	DIRECTION_UP      = 4
)

type Character struct {
	isPlayer            bool
	renderer            *sdl.Renderer
	resourceManager     *managers.ResourceManager
	CharacterInfo       *models.CharacterInfo
	Inventory           *Inventory
	message             *Message
	Pos                 *sdl.Point
	posToGo             *sdl.Point
	speed               float64
	direction           int
	currentFrame        int
	frameRate           uint32
	oldTime             uint32
	isHarvesting        bool
	resourceHarvesting  int
	harvestingFrameRate uint32
	harvestingOldTime   uint32
	building            *Building
	isInBuilding        bool
}

func NewPlayer(renderer *sdl.Renderer, resourceManager *managers.ResourceManager, characterInfo *models.CharacterInfo, x int32, y int32, isPlayer bool) *Character {
	c := &Character{
		isPlayer,
		renderer,
		resourceManager,
		characterInfo,
		NewInventory(),
		nil,
		&sdl.Point{X: x, Y: y},
		&sdl.Point{X: x, Y: y},
		5,
		DIRECTION_DEFAULT,
		0,
		100,
		0,
		false,
		-1,
		1000,
		0,
		nil,
		false,
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

func (c *Character) Build(buildingInfo *models.BuildingInfo, pos *sdl.Point) {
	c.isInBuilding = false

	newPos := &sdl.Point{}
	newPos.X = pos.X - buildingInfo.Width/2
	newPos.Y = pos.Y - buildingInfo.Height/2

	c.building = NewBuilding(c.resourceManager, buildingInfo, newPos)

	c.posToGo.X = pos.X - buildingInfo.Width/4
	c.posToGo.Y = pos.Y - buildingInfo.Height/4
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
	if c.isMoving() {
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

func (c *Character) isMoving() bool {
	if c.posToGo.X != c.Pos.X && c.posToGo.Y != c.Pos.Y {
		return true
	} else {
		return false
	}
}

func (c *Character) harvesting(mapResourceArray [][]int) {
	x := (c.Pos.X + c.GetWidth()/2) / utils.TileSize
	y := (c.Pos.Y + c.GetHeight()/2) / utils.TileSize

	if x < 1 {
		x = 1
	}

	if y < 1 {
		y = 1
	}

	if x > int32(len(mapResourceArray)-1) {
		x = int32(len(mapResourceArray) - 1)
	}

	if y > int32(len(mapResourceArray[0])-1) {
		y = int32(len(mapResourceArray[0]) - 1)
	}

	if mapResourceArray[x][y] != 0 && !c.isMoving() {
		if c.isHarvesting && c.resourceHarvesting != -1 && !c.Inventory.isResourceFull(c.resourceHarvesting) {
			if c.harvestingOldTime+c.harvestingFrameRate > sdl.GetTicks() {
				return
			}

			c.harvestingOldTime = sdl.GetTicks()
			c.Inventory.addResource(c.resourceHarvesting, 1)

			if c.Inventory.isResourceFull(c.resourceHarvesting) {
				// Show a message
				font := c.resourceManager.GetFont(utils.FONT_FIRACODE, 18)
				c.message = NewMessage("I'm tired now, I need some sleep !", font)
				c.isHarvesting = false
			}
		} else {
			c.isHarvesting = true
			c.resourceHarvesting = utils.ResourceTextureInfo[mapResourceArray[x][y]].ResourceId
			c.Inventory.addResource(c.resourceHarvesting, 0)
		}
	} else {
		c.isHarvesting = false
		c.resourceHarvesting = -1
		c.harvestingOldTime = 0
	}
}

func (c *Character) Update(MapResourceArray [][]int, camera *sdl.Rect) {
	c.move()
	c.harvesting(MapResourceArray)

	// Update the building
	if c.building != nil {
		characterPos := &sdl.Point{
			X: c.Pos.X - camera.X + c.GetWidth()/2,
			Y: c.Pos.Y - camera.Y + c.GetHeight()/2,
		}

		c.building.Update(characterPos, camera, c.isMoving())
		c.isInBuilding = c.building.IsInside(characterPos, camera, c.isMoving())
	}

	// Update messages display
	if c.message != nil {
		c.message.update()

		// Delete the message if last too long
		if c.message.doesMustBeDeleted() {
			c.message = nil
		}
	}
}

func (c *Character) Render(camera *sdl.Rect) {
	x := c.Pos.X - camera.X
	y := c.Pos.Y - camera.Y

	// Render the building
	if c.building != nil {
		characterPos := &sdl.Point{
			X: x + c.GetWidth()/2,
			Y: y + c.GetHeight()/2,
		}

		c.building.Render(c.renderer, characterPos, c.isMoving(), camera)
	}

	// Render the character
	if !c.isInBuilding && c.Pos.X+c.GetWidth() > camera.X && c.Pos.Y+c.GetHeight() > camera.Y &&
		c.Pos.X < camera.X+camera.W && c.Pos.Y < camera.Y+camera.H {
		c.renderer.Copy(
			c.resourceManager.GetTexture(c.CharacterInfo.ImageKey),
			c.getCurrentTextureRect(),
			&sdl.Rect{
				X: x,
				Y: y,
				W: c.CharacterInfo.DefaultTexture.W,
				H: c.CharacterInfo.DefaultTexture.H,
			},
		)
	}

	// Render harvesting
	if c.isHarvesting && c.isPlayer && !c.Inventory.isResourceFull(c.resourceHarvesting) {
		c.renderer.SetDrawColor(255, 255, 255, 255)
		c.renderer.DrawRect(&sdl.Rect{
			X: x,
			Y: y - HEIGHT_RECT_LIFE,
			H: HEIGHT_RECT_LIFE,
			W: c.CharacterInfo.DefaultTexture.W,
		})

		percentWidthHarvesting := float64(c.Inventory.resources[c.resourceHarvesting].amount) /
			float64(c.Inventory.resources[c.resourceHarvesting].maxAmount)

		c.renderer.FillRect(&sdl.Rect{
			X: x,
			Y: y - HEIGHT_RECT_LIFE,
			H: HEIGHT_RECT_LIFE,
			W: int32(float64(c.CharacterInfo.DefaultTexture.W) * percentWidthHarvesting),
		})
	}

	// Display the message
	if c.message != nil {
		c.message.render(x, y-HEIGHT_RECT_LIFE, c.GetWidth(), c.renderer)
	}
}
