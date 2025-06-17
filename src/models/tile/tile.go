package tile

type Tile struct {
	Properties Set
}

// makes an empty tile
func NewTile() Tile {
	return Tile{NewSet()}
}
