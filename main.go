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

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v",filename, err))
	}
	return tex 
}

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

	// Creates Player
	plr := newPlayer(renderer)

	var enemies []basicEnemy

	for i:=0; i<5; i++ {
		for j:=0; j<3; j++ {
			x := (float64(i) / 5) * screenWidth + (basicEnemySize / 2.0)
			y := float64(j) * basicEnemySize + (basicEnemySize / 2.0)

			enemy := newBasicEnemy(renderer, x, y)
			enemies = append(enemies, enemy)
		}
	}

	innitBulletPool(renderer)

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

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}

		renderer.Present()
	}

}
