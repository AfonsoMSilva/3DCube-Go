package main

import(
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

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
		sld.WINDOWPOS_UNDIFINED, sld.WINDOWPOS_UNDIFINED,
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

}