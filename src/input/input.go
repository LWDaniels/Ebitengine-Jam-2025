package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// could cache this but not convinced it's important rn

func Movement() image.Point {
	dir := image.Pt(0, 0)
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		dir.X += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		dir.X -= 1
	}
	// preventing diagonal movement
	if dir.X == 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			dir.Y -= 1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			dir.Y += 1
		}
	}
	return dir
}
