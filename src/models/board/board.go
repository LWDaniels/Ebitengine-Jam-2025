package board

import (
	"image"

	"github.com/LWDaniels/Ebitengine-Jam-2025/src/constants"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/player"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/tile"
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Layers []Layer
	Player player.Player
}

func NewBoard() Board { // should load from a text file down the line
	layers := make([]Layer, 1)
	layers[0] = NewLayer(constants.BoardWidth, constants.BoardHeight)
	return Board{layers, player.NewPlayer()}
}

func (b *Board) Draw(screen *ebiten.Image, camGeoM ebiten.GeoM) {
	b.Layers[0].Draw(screen, camGeoM)
	b.Player.Draw(screen, camGeoM)
}

func isPushAllowed(t tile.Tile, moveDir image.Point) bool {
	return (t.Properties.Contains(tile.PushableR) && moveDir.X == 1) || (t.Properties.Contains(tile.PushableD) && moveDir.Y == 1) ||
		(t.Properties.Contains(tile.PushableL) && moveDir.X == -1) || (t.Properties.Contains(tile.PushableU) && moveDir.Y == -1)
}

// tries to shift the tile at [tilePos] by [moveDir]
// returns whether that tile was successfully moved or not
func (b *Board) Move(tilePos, moveDir image.Point) bool {
	if tilePos.X < 0 || tilePos.X >= constants.BoardWidth ||
		tilePos.Y < 0 || tilePos.Y >= constants.BoardHeight {
		return false
	}
	t := b.Layers[0].Tiles[tilePos.Y][tilePos.X]
	if !t.Properties.Contains(tile.Collision) {
		return true
	}

	nextPos := tilePos.Add(moveDir)
	if isPushAllowed(t, moveDir) && b.Move(nextPos, moveDir) {
		b.Layers[0].Tiles[tilePos.Y][tilePos.X] = tile.NewTile() // may need to change if there are
		// tiles you're walking over (in fact, current system doesn't really support this at all)
		b.Layers[0].Tiles[nextPos.Y][nextPos.X] = t
		return true
	} else {
		return false
	}
}
