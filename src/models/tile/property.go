package tile

// properties either exist or don't; you can't have a property that opposes/can't be combined with another
type Property = uint

const (
	Collision Property = iota
	PushableR
	PushableD
	PushableL
	PushableU
	/*
		Other ideas:
		- Damage (maybe destroys tiles that move into it)
		- Indestructible (can't be destroyed)
		- Maybe subcategories of pushable (pushable in particular directions???)
		- Goal (not sure if the goal is to get player there or something else)
		- idk :)
	*/
)
