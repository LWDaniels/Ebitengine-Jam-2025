package tile

import (
	"image"
	"math"

	"github.com/LWDaniels/Ebitengine-Jam-2025/assets"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Properties Set
}

// makes an empty tile
func NewTile() Tile {
	return Tile{NewSet()}
}

func (t Tile) Draw(screen *ebiten.Image, camGeoM ebiten.GeoM, tilePos image.Point) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(constants.Yellow)
	op.GeoM.Translate(float64(tilePos.X*constants.TileSize), float64(tilePos.Y*constants.TileSize))
	op.GeoM.Concat(camGeoM)
	screen.DrawImage(assets.Block, op)

	pushableIm := assets.Arrow
	arrowGeoM := ebiten.GeoM{}
	arrowGeoM.Translate(-float64(pushableIm.Bounds().Dx()), -float64(pushableIm.Bounds().Dy())/2)
	for moveProp := PushableR; moveProp <= PushableU; moveProp++ {
		if t.Properties.Contains(moveProp) {
			opArrow := &ebiten.DrawImageOptions{}
			opArrow.ColorScale.ScaleWithColor(constants.Teal)
			opArrow.GeoM.Translate(constants.TileSize/2, 0)
			opArrow.GeoM.Concat(arrowGeoM)
			opArrow.GeoM.Translate(constants.TileSize/2, constants.TileSize/2)
			opArrow.GeoM.Concat(op.GeoM)
			screen.DrawImage(pushableIm, opArrow)
		}
		arrowGeoM.Rotate(math.Pi / 2)
	}
}
