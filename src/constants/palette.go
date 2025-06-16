package constants

import (
	"image/color"
	"strconv"
)

// https://colorhunt.co/palette/4da8da80d8c3ffd66bf5f5f5
const (
	blueHex     = "4DA8DA"
	tealHex     = "80D8C3"
	yellowHex   = "FFD66B"
	offWhiteHex = "F5F5F5"
)

var (
	Blue     = toColor(blueHex)
	Teal     = toColor(tealHex)
	Yellow   = toColor(yellowHex)
	OffWhite = toColor(offWhiteHex)
)

func toColor(hex string) color.Color {
	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}
