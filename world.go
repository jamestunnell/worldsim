package worldsim

type World interface {
	Agents() map[string]Agent
	Objects() map[string]Object
	Areas() map[string]Area
}
