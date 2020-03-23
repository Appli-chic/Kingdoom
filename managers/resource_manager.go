package managers

import (
	"github.com/veandco/go-sdl2/img"
	"fmt"
	"os"
	"github.com/veandco/go-sdl2/sdl"
)

type ResourceManager struct {
	images   map[int]*sdl.Surface
	textures map[int]*sdl.Texture
}

func NewResourceManager() ResourceManager {
	r := ResourceManager{
		map[int]*sdl.Surface{},
		map[int]*sdl.Texture{},
	}

	return r
}

func (r *ResourceManager) GetImage(key int) *sdl.Surface {
	return r.images[key]
}

func (r *ResourceManager) GetTexture(key int) *sdl.Texture {
	return r.textures[key]
}

func (r *ResourceManager) LoadImage(key int) error {
	image, err := img.Load("assets/tileset/outside2.png")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load image: %s\n", err)
	} else {
		r.images[key] = image
	}

	return err
}

func (r *ResourceManager) LoadTextureFromImage(key int, image *sdl.Surface, renderer *sdl.Renderer) error {
	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
	} else {
		r.textures[key] = texture
	}

	return err
}
