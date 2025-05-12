package facade

// HomeCinema is the facade for controlling all of the devices of home theater in the correct order
type HomeCinema struct {
	projector *Projector
	server    *MediaServer
	lights    *Lights
	sound     *SoundSystem
}

// Prepare allows to prepare home theater for the projection
func (c *HomeCinema) Prepare() {
	c.lights.On()
	c.sound.On()
	c.sound.SetVolume(1)
	c.projector.On()
	c.server.On()
	c.projector.SetInput(c.server.OutputInterface)
}

// PlayMovie allows the devices to play the actual movie
func (c *HomeCinema) PlayMovie(name string) {
	c.sound.SetVolume(60)
	c.lights.SetBrightness(0)
	c.server.Play(name)
}

// EndMovie allows to stop the projection of the movie
func (c *HomeCinema) EndMovie() {
	c.lights.SetBrightness(100)
	c.server.Stop()
}

// Close allows to turn off the devices and close the home theater until the next projection
func (c *HomeCinema) Close() {
	c.sound.Off()
	c.projector.Off()
	c.server.Off()
	c.lights.Off()
}

// NewHomeTheater is the factory function to create the new home theater facade
func NewHomeTheater(input uint8, projectorName, serverName, lightsName, soundSystemName string) *HomeCinema {
	return &HomeCinema{
		projector: &Projector{name: projectorName},
		server:    &MediaServer{name: serverName, OutputInterface: input},
		lights:    &Lights{description: lightsName},
		sound:     &SoundSystem{name: soundSystemName},
	}
}
