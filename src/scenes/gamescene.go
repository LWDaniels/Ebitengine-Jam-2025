package scenes

import (
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/constants"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/input"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/board"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/camera"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/tile"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	cam   camera.Camera
	board board.Board
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(constants.OffWhite)
	g.board.Draw(screen, g.cam.GeoM())
}

func (g *GameScene) Update() error {
	moveDir := input.Movement()
	nextPos := g.board.Player.Pos.Add(moveDir)
	if g.board.Move(nextPos, moveDir) {
		g.board.Player.Pos = nextPos
	}

	return nil
}

func (g *GameScene) Layout(screenWidth, screenHeight int) (targetWidth, targetHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func NewGameScene() ebiten.Game {
	game := &GameScene{camera.NewCamera(0, 0, 4), board.NewBoard()}
	t := tile.NewTile()
	t.Properties.Add(tile.Collision)
	t.Properties.Add(tile.Pushable)
	game.board.Layers[0].Tiles[5][5] = t

	// t2 := tile.NewTile()
	// t2.Properties.Add(tile.Collision)
	// game.board.Layers[0].Tiles[2][7] = t2
	return game
}
