package camera

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	pos   image.Point
	scale int // maybe this should be a float??? idk, for pixel art it should be easily divisible or something
}

func NewCamera(x, y, scale int) Camera {
	return Camera{image.Pt(x, y), scale}
}

func (cam Camera) GeoM() ebiten.GeoM {
	geoM := ebiten.GeoM{}
	geoM.Scale(float64(cam.scale), float64(cam.scale))
	geoM.Translate(float64(cam.pos.X), float64(cam.pos.Y))
	return geoM
}
