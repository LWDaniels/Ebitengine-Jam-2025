package main

import (
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Jamming")

	scene := scenes.NewGameScene()
	if err := ebiten.RunGame(scene); err != nil {
		panic(err)
	}
}
