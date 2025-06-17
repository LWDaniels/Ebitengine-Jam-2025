package board

import (
	"image"

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

			l.Tiles[y][x].Draw(screen, camGeoM, image.Pt(x, y))
		}
	}
}
