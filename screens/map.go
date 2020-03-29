package screens

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const NB_BIOMES = 3
const NB_INT_BIOME = 20

type Map struct {
	MapArray [][]int
}

func NewMap(width int, height int) *Map {
	matrix := make([][]int, width)
	rows := make([]int, width*height)
	for i := 0; i < width; i++ {
		matrix[i] = rows[i*height : (i+1)*height]
	}

	m := &Map{
		matrix,
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

func (m *Map) roundBorders() {
	for x := 1; x < len(m.MapArray)-1; x++ {
		for y := 1; y < len(m.MapArray[x])-1; y++ {
			m.roundBorderBetweenTwoBiomes(x, y, utils.DIRT, utils.PLAIN)
			m.roundBorderBetweenTwoBiomes(x, y, utils.WATER, utils.PLAIN)
		}
	}
}

func (m *Map) createRiver() {
	x := 1.
	y := 5.
	endRiverX := 100.
	endRiverY := 100.

	m.MapArray[int(x)][int(y)] = utils.WATER

	for {
		finalResults := []float64{}

		// Calculate Top
		distanceToEnd := math.Abs(x-endRiverX) + math.Abs(y-endRiverY)
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
		} else {
			// There is a second direction possible
			randomVal := rand.Intn(2)

			if randomVal == 0 {
				finalDirection = direction
			} else {
				finalDirection = secondDirection
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
		if int(x) < len(m.MapArray)-1 && int(y) < len(m.MapArray[0])-1 && x > 0 && y > 0 {
			m.MapArray[int(x)][int(y)] = utils.WATER
		}

		// We found the end of the river
		if x == endRiverX && y == endRiverY {
			break
		}
	}
}

func (m *Map) initMap(width int, height int) {
	// Create rivers
	m.createRiver()

	// Create the map
	// for x := 0; x < len(m.MapArray); x++ {
	// 	for y := 0; y < len(m.MapArray[x]); y++ {
	// 	}
	// }

	m.roundBorders()
}

func (m *Map) displaysTile(camera *sdl.Rect, resourceManager *managers.ResourceManager,
	renderer *sdl.Renderer, x int, y int) {
	tileInfo := utils.GroundTextureInfo[m.MapArray[x][y]]

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

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			m.displaysTile(camera, resourceManager, renderer, x, y)
		}
	}
}
