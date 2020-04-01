package managers

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
)

type AudioManager struct {
	actualMusic *mix.Music
	loopsMusic  int
}

func NewAudioManager() *AudioManager {
	return &AudioManager{}
}

func (a *AudioManager) PlayMusic(music *mix.Music, loops int) {
	err := music.Play(1)

	if err != nil {
		log.Println(err)
	} else {
		a.actualMusic = music
		a.loopsMusic = loops
	}
}
