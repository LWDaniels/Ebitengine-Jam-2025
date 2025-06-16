package camera

import (
	"image"

	"github.com/LWDaniels/Ebitengine-Jam-2025/src/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Pos   image.Point
	Scale int // maybe this should be a float??? idk, for pixel art it should be easily divisible or something
}

func NewCamera(x, y, scale int) Camera {
	return Camera{image.Pt(x, y), scale}
}

func (cam Camera) GeoM() ebiten.GeoM {
	geoM := ebiten.GeoM{}
	geoM.Scale(float64(cam.Scale), float64(cam.Scale))
	geoM.Translate(float64(cam.Pos.X), float64(cam.Pos.Y))
	return geoM
}

func (cam Camera) Bounds() image.Rectangle {
	minX := cam.Pos.X / cam.Scale
	minY := cam.Pos.Y / cam.Scale
	maxX := minX + constants.ScreenWidth/cam.Scale
	maxY := minY + constants.ScreenHeight/cam.Scale
	return image.Rect(minX, minY, maxX, maxY)
}
