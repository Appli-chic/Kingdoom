package screens

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const NB_BIOMES = 3
const NB_INT_BIOME = 20

type BiomeInfo struct {
	key int
	x   int
	y   int
}

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

func (m *Map) createSeed(width int, height int, params ...int) *BiomeInfo {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(width - 1)

	rand.Seed(time.Now().UnixNano())
	y := rand.Intn(height - 1)

	rand.Seed(time.Now().UnixNano())
	biomeIndex := rand.Intn(NB_BIOMES)

	var seed int

	if params == nil {
		seed = utils.PLAIN

		switch biomeIndex {
		case 0:
			seed = utils.PLAIN
		case 1:
			seed = utils.DIRT
		case 2:
			seed = utils.WATER
		default:
			seed = utils.PLAIN
		}
	} else {
		seed = params[0]
	}

	return &BiomeInfo{
		key: seed,
		x:   x,
		y:   y,
	}
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

func (m *Map) initMap(width int, height int) {
	biomeInfoList := []*BiomeInfo{}

	// Create plains
	for i := 0; i < int(float64(width)/2.5); i++ {
		seedBiomeInfo := m.createSeed(width, height, utils.PLAIN)
		biomeInfoList = append(biomeInfoList, seedBiomeInfo)
	}

	// Create deserts
	for i := 0; i < int(float64(width)/10); i++ {
		seedBiomeInfo := m.createSeed(width, height, utils.DIRT)
		biomeInfoList = append(biomeInfoList, seedBiomeInfo)
	}

	// Create lakes
	for i := 0; i < int(float64(width)/10); i++ {
		seedBiomeInfo := m.createSeed(width, height, utils.WATER)
		biomeInfoList = append(biomeInfoList, seedBiomeInfo)
	}

	// Create the map
	for x := 0; x < len(m.MapArray); x++ {
		for y := 0; y < len(m.MapArray[x]); y++ {
			nearest := -1
			dist := 99999999

			for i := 0; i < len(biomeInfoList); i++ {
				xdiff := biomeInfoList[i].x - x
				ydiff := biomeInfoList[i].y - y

				cdist := xdiff*xdiff + ydiff*ydiff

				if cdist < dist {
					nearest = biomeInfoList[i].key
					dist = cdist
				}
			}

			m.MapArray[x][y] = nearest
		}
	}

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
