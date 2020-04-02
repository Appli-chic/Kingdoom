package screens

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const NB_BIOMES = 3
const NB_INT_BIOME = 20
const NB_INT_WATER_ANIMATION = 1000

type Map struct {
	MapArray          [][]int
	MapElementArray   [][]int
	currentFrameWater int
	frameRateWater    uint32
	oldTimeWater      uint32
}

func NewMap(width int, height int) *Map {
	matrix := make([][]int, width)
	rows := make([]int, width*height)
	for i := 0; i < width; i++ {
		matrix[i] = rows[i*height : (i+1)*height]
	}

	matrixElement := make([][]int, width)
	rowsElement := make([]int, width*height)
	for i := 0; i < width; i++ {
		matrixElement[i] = rowsElement[i*height : (i+1)*height]
	}

	m := &Map{
		matrix,
		matrixElement,
		0,
		150,
		0,
	}

	m.initMap(width, height)

	return m
}

func (m *Map) roundBorderBetweenTwoBiomes(x int, y int, biome1 int, biome2 int) {
	if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y-1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Corner Left Up
		m.MapArray[x][y] = biome1 + 9
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y-1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Corner Right Up
		m.MapArray[x][y] = biome1 + 10
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y+1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Corner Left Bottom
		m.MapArray[x][y] = biome1 + 11
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y+1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Corner Right Bottom
		m.MapArray[x][y] = biome1 + 12
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME == biome2/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Left Up
		m.MapArray[x][y] = biome1 + 5
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME == biome2/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Right Up
		m.MapArray[x][y] = biome1 + 6
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME == biome2/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Left Bottom
		m.MapArray[x][y] = biome1 + 7
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME == biome2/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Right Bottom
		m.MapArray[x][y] = biome1 + 8
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Left
		m.MapArray[x][y] = biome1 + 1
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Right
		m.MapArray[x][y] = biome1 + 2
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x][y+1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Down
		m.MapArray[x][y] = biome1 + 3
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x][y-1]/NB_INT_BIOME == biome2/NB_INT_BIOME {
		// Up
		m.MapArray[x][y] = biome1 + 4
	}
}

func (m *Map) roundBorderOneBiomeAgainstAll(x int, y int, biome1 int) {
	if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y-1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Corner Left Up
		m.MapArray[x][y] = biome1 + 9
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y-1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Corner Right Up
		m.MapArray[x][y] = biome1 + 10
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y+1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Corner Left Bottom
		m.MapArray[x][y] = biome1 + 11
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME == biome1/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y+1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Corner Right Bottom
		m.MapArray[x][y] = biome1 + 12
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME != biome1/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Left Up
		m.MapArray[x][y] = biome1 + 5
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME != biome1/NB_INT_BIOME &&
		m.MapArray[x][y-1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Right Up
		m.MapArray[x][y] = biome1 + 6
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME != biome1/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Left Bottom
		m.MapArray[x][y] = biome1 + 7
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME != biome1/NB_INT_BIOME &&
		m.MapArray[x][y+1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Right Bottom
		m.MapArray[x][y] = biome1 + 8
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x-1][y]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Left
		m.MapArray[x][y] = biome1 + 1
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x+1][y]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Right
		m.MapArray[x][y] = biome1 + 2
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x][y+1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Down
		m.MapArray[x][y] = biome1 + 3
	} else if m.MapArray[x][y]/NB_INT_BIOME == biome1/NB_INT_BIOME && m.MapArray[x][y-1]/NB_INT_BIOME != biome1/NB_INT_BIOME {
		// Up
		m.MapArray[x][y] = biome1 + 4
	}
}

func (m *Map) roundBorders() {
	for x := 1; x < len(m.MapArray)-1; x++ {
		for y := 1; y < len(m.MapArray[x])-1; y++ {
			m.roundBorderOneBiomeAgainstAll(x, y, utils.DIRT)
			m.roundBorderOneBiomeAgainstAll(x, y, utils.WATER)
		}
	}
}

func (m *Map) growRiverSize(x int, y int, size int, lastDirection int) {
	for i := -size; i < size; i++ {
		if lastDirection == 0 || lastDirection == 2 {
			if x+i < len(m.MapArray) && y < len(m.MapArray[0]) && x+i >= 0 && y >= 0 {
				m.MapArray[x+i][y] = utils.WATER
			}
		} else {
			if x < len(m.MapArray) && y+i < len(m.MapArray[0]) && x >= 0 && y+i >= 0 {
				m.MapArray[x][y+i] = utils.WATER
			}
		}
	}
}

func (m *Map) createRiver() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(10-3) + 3

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	x := 0.
	y := float64(r.Intn(len(m.MapArray)))

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	endRiverX := float64(len(m.MapArray) - 1)
	endRiverY := float64(r.Intn(len(m.MapArray[0])))
	lastDirection := -1

	m.MapArray[int(x)][int(y)] = utils.WATER
	m.growRiverSize(0, int(y), size, 3)

	for {
		finalResults := []float64{}

		// Calculate Top
		distanceToEnd := math.Abs(x-endRiverX) + math.Abs(y-1-endRiverY)
		finalResults = append(finalResults, distanceToEnd)

		// Calculate Right
		distanceToEnd = math.Abs(x+1-endRiverX) + math.Abs(y-endRiverY)
		finalResults = append(finalResults, distanceToEnd)

		// Calculate Bottom
		distanceToEnd = math.Abs(x-endRiverX) + math.Abs(y+1-endRiverY)
		finalResults = append(finalResults, distanceToEnd)

		// Calculate Left
		distanceToEnd = math.Abs(x-1-endRiverX) + math.Abs(y-endRiverY)
		finalResults = append(finalResults, distanceToEnd)

		// Choose which direction to take
		smallestNumber := 999999999999.
		direction := -1
		secondDirection := -1
		finalDirection := -1

		for i := 0; i < len(finalResults); i++ {
			if finalResults[i] < smallestNumber {
				smallestNumber = finalResults[i]
				direction = i
			}
		}

		for i := 0; i < len(finalResults); i++ {
			if finalResults[i] == smallestNumber && i != direction {
				secondDirection = i
			}
		}

		if secondDirection < 0 {
			// There is only one direction
			finalDirection = direction
			lastDirection = direction
		} else {
			// There is a second direction possible
			r = rand.New(rand.NewSource(time.Now().UnixNano()))
			randomVal := r.Intn(2)

			if randomVal == 0 {
				finalDirection = direction
				lastDirection = direction
			} else {
				finalDirection = secondDirection
				lastDirection = secondDirection
			}
		}

		switch finalDirection {
		case 0:
			y -= 1 // Top
			break
		case 1:
			x += 1 // Right
			break
		case 2:
			y += 1 // Bottom
			break
		case 3:
			x -= 1 // Left
			break
		}

		// Set the new tile as water
		m.MapArray[int(x)][int(y)] = utils.WATER
		m.growRiverSize(int(x), int(y), size, lastDirection)

		// We found the end of the river
		if x == endRiverX && y == endRiverY {
			break
		}
	}
}

func (m *Map) generateBiome(biome int, width int, height int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sizeX := r.Intn((width/3)-20) + 20

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	sizeY := r.Intn((height/3)-20) + 20

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	seedX := r.Intn(width-sizeX) + sizeX

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	seedY := r.Intn(height-sizeY) + sizeY

	for x := seedX - sizeX; x < seedX+sizeX; x++ {
		for y := seedY - sizeY; y < seedY+sizeY; y++ {
			if x < len(m.MapArray) && y < len(m.MapArray[0]) && x >= 0 && y >= 0 {
				m.MapArray[x][y] = biome
			}
		}
	}
}

func (m *Map) addElementToMap(element int, width int, height int) {
	isPlaceAvailable := false

	for !isPlaceAvailable {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := r.Intn(width)

		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		y := r.Intn(height)

		if m.MapArray[x][y] != utils.WATER && m.MapElementArray[x][y] == 0 {
			m.MapElementArray[x][y] = element
			isPlaceAvailable = true
		}
	}
}

func (m *Map) initMap(width int, height int) {
	// Create biome
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberDirtBiome := r.Intn(4-1) + 1

	for i := 0; i < numberDirtBiome; i++ {
		m.generateBiome(utils.DIRT, width, height)
	}

	// Create river
	m.createRiver()

	// Add resources
	nbTrees := width / 10

	for i := 0; i < nbTrees; i++ {
		m.addElementToMap(utils.TREE, width, height)
	}

	m.MapElementArray[10][10] = utils.TREE

	// Round borders
	m.roundBorders()
}

func (m *Map) animateWater() {
	if m.oldTimeWater+m.frameRateWater > sdl.GetTicks() {
		return
	}

	m.oldTimeWater = sdl.GetTicks()
	m.currentFrameWater++

	if m.currentFrameWater > 2 {
		m.currentFrameWater = 0
	}
}

func (m *Map) Update() {
	m.animateWater()
}

func (m *Map) displaysTile(camera *sdl.Rect, resourceManager *managers.ResourceManager,
	renderer *sdl.Renderer, x int, y int) {
	tileInfo := utils.GroundTextureInfo[m.MapArray[x][y]]

	// Take the animating water tile
	if m.MapArray[x][y]/NB_INT_BIOME == utils.WATER/NB_INT_BIOME {
		tileInfo = utils.GroundTextureInfo[m.MapArray[x][y]+m.currentFrameWater*NB_INT_WATER_ANIMATION]
	}

	err := renderer.Copy(
		resourceManager.GetTexture(tileInfo.ImageKey),
		tileInfo.Src,
		&sdl.Rect{
			X: int32(TileSize*x) - camera.X,
			Y: int32(TileSize*y) - camera.Y,
			W: TileSize,
			H: TileSize,
		},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to copy: %s\n", err)
	}
}

func (m *Map) displaysElement(camera *sdl.Rect, resourceManager *managers.ResourceManager,
	renderer *sdl.Renderer, x int, y int) {
	if m.MapElementArray[x][y] != 0 {
		tileInfo := utils.GroundTextureInfo[m.MapElementArray[x][y]]

		err := renderer.Copy(
			resourceManager.GetTexture(tileInfo.ImageKey),
			tileInfo.Src,
			&sdl.Rect{
				X: int32(TileSize*x) - camera.X,
				Y: int32(TileSize*y) - camera.Y,
				W: tileInfo.Src.W,
				H: tileInfo.Src.H,
			},
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to copy: %s\n", err)
		}
	}
}

func (m *Map) Render(camera *sdl.Rect, resourceManager *managers.ResourceManager,
	renderer *sdl.Renderer) {
	minX := int(camera.X/TileSize) - 2
	maxX := int((camera.X+camera.W)/TileSize) + 2
	minY := int(camera.Y/TileSize) - 2
	maxY := int((camera.Y+camera.H)/TileSize) + 2

	if minX < 1 {
		minX = 1
	}

	if minY < 1 {
		minY = 1
	}

	if maxX > len(m.MapArray)-1 {
		maxX = len(m.MapArray) - 1
	}

	if maxY > len(m.MapArray[0])-1 {
		maxY = len(m.MapArray[0]) - 1
	}

	// Displays ground
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			m.displaysTile(camera, resourceManager, renderer, x, y)
		}
	}

	// Displays elements
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			m.displaysElement(camera, resourceManager, renderer, x, y)
		}
	}

}
