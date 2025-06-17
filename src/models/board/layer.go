package board

import (
	"github.com/LWDaniels/Ebitengine-Jam-2025/assets"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/constants"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/tile"
	"github.com/hajimehoshi/ebiten/v2"
)

type Layer struct {
	Tiles [][]tile.Tile
}

// returns an empty layer of the correct size
func NewLayer(width, height int) Layer {
	l := Layer{}
	l.Tiles = make([][]tile.Tile, height)
	for y := range l.Tiles {
		l.Tiles[y] = make([]tile.Tile, width)
	}
	return l
}

func (l Layer) Draw(screen *ebiten.Image, camGeoM ebiten.GeoM) {
	for y := range l.Tiles {
		for x := range l.Tiles[y] {
			if l.Tiles[y][x].Properties.Size() == 0 {
				// empty
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.ColorScale.ScaleWithColor(constants.Yellow)
			op.GeoM.Translate(float64(x*constants.TileSize), float64(y*constants.TileSize))
			op.GeoM.Concat(camGeoM)
			screen.DrawImage(assets.Block, op)
		}
	}
}
