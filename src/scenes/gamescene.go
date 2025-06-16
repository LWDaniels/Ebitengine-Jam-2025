package scenes

import (
	"github.com/LWDaniels/Ebitengine-Jam-2025/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(assets.Gopher, &ebiten.DrawImageOptions{})
}

func (g *GameScene) Update() error {
	return nil
}

func (g *GameScene) Layout(screenWidth, screenHeight int) (targetWidth, targetHeight int) {
	return 1280, 720
}

func NewGameScene() ebiten.Game {
	return &GameScene{}
}
