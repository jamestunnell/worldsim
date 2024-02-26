package worldsim

type Agent interface {
	Identifyable
	Locatable

	Sense()
	Act()
}
