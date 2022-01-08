package main

import(
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

// Window Size
const (
	screenWidth = 600
	screenHeight = 800
)

func main() {

	// Initialization
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initialization Error: ", err)
		return
	}

	// Creating Window
	window, err := sdl.CreateWindow(
		"3DCube",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("CreateWindow Error: ", err)
		return
	}
	defer window.Destroy()

	// Creating Renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Renderer Error: ", err)
		return
	}
	defer renderer.Destroy()

	plr, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("creating player:", err)
		return
	}

	enemy, err := newBasicEnemy(renderer, screenWidth/2.0, screenHeight/2.0)
	if err != nil {
		fmt.Println("creating basicEnemy:", err)
		return
	}

	for{

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		plr.draw(renderer)
		plr.update()
		enemy.draw(renderer)
		renderer.Present()
	}

}
