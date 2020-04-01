package managers

import (
	"fmt"
	"os"

	"github.com/kingdoom/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type ResourceManager struct {
	textures map[int]*sdl.Texture
	audios   map[int]*mix.Music
}

func NewResourceManager() ResourceManager {
	r := ResourceManager{
		map[int]*sdl.Texture{},
		map[int]*mix.Music{},
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

func (r *ResourceManager) GetAudio(key int) *mix.Music {
	return r.audios[key]
}

func (r *ResourceManager) LoadAudio(key int) {
	music, err := mix.LoadMUS(utils.AudioPath[key])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load audio: %s\n", err)
	} else {
		r.audios[key] = music
	}
}
