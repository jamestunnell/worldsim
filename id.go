package worldsim

import (
	"github.com/jaevor/go-nanoid"
	"github.com/rs/zerolog/log"
)

type Identifyable interface {
	ID() string
}

// The canonic NanoID is nanoid.Standard(21).
var idGenerator func() string

func init() {
	// canonical ID generator
	g, err := nanoid.Standard(21)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to make ID generator")
	}

	idGenerator = g
}

func NewID() string {
	return idGenerator()
}
