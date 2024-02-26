package worldsim

type Locatable interface {
	AreaID() string
	AreaCoords() []float64
}
