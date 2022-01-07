package main

import(
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

// Window Size
const (
	screenWidth = 800
	screenHeight = 600
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

	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil{
		fmt.Println("loading the player:", err)
		return
	}
	defer img.Free()

	playerTex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("creating texture:")
	}
	defer playerTex.Destroy()

	for{

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Copy(playerTex,
			&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
			&sdl.Rect{X: 0, Y: 0, W: 105, H: 105})

		renderer.Present()
	}

}
