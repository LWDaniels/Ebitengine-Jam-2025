package scenes

import (
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/constants"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/input"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/board"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/models/camera"
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
	g.board.Player.Pos = g.board.Player.Pos.Add(moveDir)

	return nil
}

func (g *GameScene) Layout(screenWidth, screenHeight int) (targetWidth, targetHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func NewGameScene() ebiten.Game {
	game := &GameScene{camera.NewCamera(0, 0, 4), board.NewBoard()}
	return game
}
