package player

import (
	"image"

	"github.com/LWDaniels/Ebitengine-Jam-2025/assets"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewPlayer() Player {
	tint := ebiten.ColorScale{}
	tint.ScaleWithColor(constants.Teal)
	return Player{image.Pt(0, 0), assets.Block, tint}
}

type Player struct {
	Pos   image.Point
	Image *ebiten.Image
	Tint  ebiten.ColorScale
}

func (p Player) Draw(screen *ebiten.Image, camGeoM ebiten.GeoM) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale = p.Tint
	op.GeoM.Translate(float64(p.Pos.X*constants.TileSize), float64(p.Pos.Y*constants.TileSize))
	op.GeoM.Concat(camGeoM)
	screen.DrawImage(p.Image, op)
}
