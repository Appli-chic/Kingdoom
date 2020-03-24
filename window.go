package main

import (
	"fmt"
	"github.com/kingdoom/screens"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Window struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	screen   screens.Screen
}

func NewWindow() *Window {
	w := &Window{}
	return w
}

func (w *Window) Show(title string, width int32, height int32) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}
	defer renderer.Destroy()

	w.window = window
	w.renderer = renderer
	w.screen = screens.NewGameScreen(window, renderer)

	var fps uint32
	fps = 60
	var delay uint32
	delay = 1000 / fps

	running := true
	for running {
		frameStart := sdl.GetTicks()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}

			if running {
				running = w.screen.HandleEvents(event)
			}
		}

		w.screen.Update()
		w.screen.Render()

		var frameTime uint32
		frameTime = sdl.GetTicks() - frameStart
		if frameTime < delay {
			sdl.Delay(delay - frameTime)
		}
	}
}
