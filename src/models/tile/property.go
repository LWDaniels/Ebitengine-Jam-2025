package tile

// properties either exist or don't; you can't have a property that opposes/can't be combined with another
type Property = uint

const (
	Collision = iota
	Pushable
)
