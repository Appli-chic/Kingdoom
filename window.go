package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/kingdoom/screens"
	"fmt"
	"os"
)

type Window struct {
	window  *sdl.Window
	renderer *sdl.Renderer
	screen  screens.Screen
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

	running := true
	for running {
		running = w.screen.HandleEvents()
		w.screen.Update()
		w.screen.Render()
	}
}
