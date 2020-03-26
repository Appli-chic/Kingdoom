package screens

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aquilax/go-perlin"
	"github.com/kingdoom/managers"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type Map struct {
	MapArray [][]int
	alpha    float64
	beta     float64
	n        int
	seed     int64
}

func NewMap(width int, height int) *Map {
	matrix := make([][]int, width)
	rows := make([]int, width*height)
	for i := 0; i < width; i++ {
		matrix[i] = rows[i*height : (i+1)*height]
	}

	// Defines the number of seeds
	rand.Seed(time.Now().UnixNano())
	numberSeeds := rand.Int63()

	m := &Map{
		matrix,
		2.,
		2.,
		9,
		numberSeeds,
	}

	m.initMap(width, height)

	return m
}

func (m *Map) initMap(width int, height int) {
	// Create transition map with perlin noise
	matrix := make([][]float64, width)
	rows := make([]float64, width*height)
	for i := 0; i < width; i++ {
		matrix[i] = rows[i*height : (i+1)*height]
	}

	// Show the data
	p := perlin.NewPerlin(m.alpha, m.beta, m.n, m.seed)
	for x := 0.; x < float64(len(matrix)); x++ {
		for y := 0.; y < float64(len(matrix[int(x)])); y++ {
			matrix[int(x)][int(y)] = p.Noise2D(x/10, y/10)
		}
	}

	// Create real map
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if matrix[x][y] >= 0.3 {
				m.MapArray[x][y] = utils.WATER_GRASS
			} else if matrix[x][y] >= -0.4 {
				m.MapArray[x][y] = utils.PLAIN
			} else {
				m.MapArray[x][y] = utils.DIRT
			}
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
