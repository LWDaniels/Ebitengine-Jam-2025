package assets

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png" // to enable png decoding

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed gopher.png
	gopher_png []byte
	Gopher     = newTexture(gopher_png)
	//go:embed simple_block.png
	block_png []byte
	Block     = newTexture(block_png)
)

func newTexture(byteArr []byte) *ebiten.Image {
	im, _, _ := image.Decode(bytes.NewReader(byteArr))
	return ebiten.NewImageFromImage(im)
}
