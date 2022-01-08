package main

import (
"github.com/veandco/go-sdl2/sdl"
"fmt"
)

const basicEnemySize = 105

type basicEnemy struct {
	tex *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy, err error){
	img, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return basicEnemy{}, fmt.Errorf("loading the basic_enemy: %v", err)
	}
	defer img.Free()

	be.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return basicEnemy{}, fmt.Errorf("creating basic_enemy texture: %v", err)
	}
	be.x = x
	be.y = y
	return be, nil
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	x := be.x - basicEnemySize/2.0
	y := be.y - basicEnemySize/2.0

	renderer.CopyEx(be.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
		180,
		sdl.Point{X: basicEnemySize/2.0, Y: basicEnemySize/2.0})
}
