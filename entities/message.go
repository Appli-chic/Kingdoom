package entities

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Message struct {
	font       *ttf.Font
	text       string
	duration   uint32
	oldTime    uint32
	currenTime uint32
	frameRate  uint32
	fontSize   int
}

func NewMessage(text string, font *ttf.Font) *Message {
	fontSize := 18

	return &Message{font, text, 5000, 0, 0, 1000, fontSize}
}

func (m *Message) doesMustBeDeleted() bool {
	return m.currenTime >= m.duration
}

func (m *Message) update() {
	if m.oldTime+m.frameRate > sdl.GetTicks() {
		return
	}

	m.oldTime = sdl.GetTicks()
	m.currenTime += m.frameRate
}

func (m *Message) render(x int32, y int32, width int32, renderer *sdl.Renderer) {
	solid, err := m.font.RenderUTF8Solid(m.text, sdl.Color{255, 255, 255, 255})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render text: %s\n", err)
		return
	}

	defer solid.Free()

	texture, err := renderer.CreateTextureFromSurface(solid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render create text texture: %s\n", err)
		return
	}

	renderer.Copy(
		texture,
		nil,
		&sdl.Rect{x - (solid.ClipRect.W / 2) + width/2, y - 8, solid.ClipRect.W, solid.ClipRect.H},
	)
}
