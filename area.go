package worldsim

type Area interface {
	Identifyable

	Neighbors()
}
