package facade

// Switcher is the general interface for all the devices of Home Theater
type Switcher interface {
	On()
	Off()
}

// Player is the general interface for all the devices able to play the movies
type Player interface {
	Play(string)
	Stop()
}
