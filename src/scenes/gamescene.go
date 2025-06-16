package scenes

import (
	"github.com/LWDaniels/Ebitengine-Jam-2025/assets"
	"github.com/LWDaniels/Ebitengine-Jam-2025/src/scenes/camera"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	cam camera.Camera
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Concat(g.cam.GeoM())
	screen.DrawImage(assets.Gopher, op)
}

func (g *GameScene) Update() error {
	return nil
}

func (g *GameScene) Layout(screenWidth, screenHeight int) (targetWidth, targetHeight int) {
	return 1280, 720
}

func NewGameScene() ebiten.Game {
	return &GameScene{camera.NewCamera(0, 0, 1)}
}
