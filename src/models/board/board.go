package board

import (
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/player"
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Layers []Layer
	Player player.Player
}

func NewBoard() Board { // should load from a text file down the line
	return Board{nil, player.NewPlayer()}
}

func (b Board) Draw(screen *ebiten.Image, camGeoM ebiten.GeoM) {
	b.Player.Draw(screen, camGeoM)
}
