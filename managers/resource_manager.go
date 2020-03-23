package managers

import (
	"fmt"
	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type ResourceManager struct {
	textures map[int]*sdl.Texture
}

func NewResourceManager() ResourceManager {
	r := ResourceManager{
		map[int]*sdl.Texture{},
	}

	return r
}

func (r *ResourceManager) GetTexture(key int) *sdl.Texture {
	return r.textures[key]
}

func (r *ResourceManager) LoadTexture(key int, renderer *sdl.Renderer) {
	image, err := img.Load(utils.ImagesPath[key])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load image: %s\n", err)
	}

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
	} else {
		r.textures[key] = texture
	}
}
