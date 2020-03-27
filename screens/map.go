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

const NB_BIOMES = 4

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
			seed = utils.SAND
		case 3:
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

func (m *Map) initMap(width int, height int) {
	biomeInfoList := []*BiomeInfo{}

	// Create plains
	for i := 0; i < 100; i++ {
		seedBiomeInfo := m.createSeed(width, height, utils.PLAIN)
		biomeInfoList = append(biomeInfoList, seedBiomeInfo)
	}

	// Create deserts
	for i := 0; i < 10; i++ {
		seedBiomeInfo := m.createSeed(width, height, utils.SAND)
		biomeInfoList = append(biomeInfoList, seedBiomeInfo)
	}

	// Create rivers
	for i := 0; i < 10; i++ {
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

	if minX < 0 {
		minX = 0
	}

	if minY < 0 {
		minY = 0
	}

	if maxX > len(m.MapArray) {
		maxX = len(m.MapArray)
	}

	if maxY > len(m.MapArray[0]) {
		maxY = len(m.MapArray[0])
	}

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			m.displaysTile(camera, resourceManager, renderer, x, y)
		}
	}
}
