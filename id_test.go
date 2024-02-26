package worldsim_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/worldsim"
)

func TestNewID_IdsAreNonEmptyAndUnique(t *testing.T) {
	id1 := worldsim.NewID()
	id2 := worldsim.NewID()

	assert.NotEmpty(t, id1)
	assert.Equal(t, len(id1), len(id2))
	assert.NotEqual(t, id1, id2)
}

func TestNewID_Supports10KUniqueIDs(t *testing.T) {
	ids := make([]string, 10000)

	for i := 0; i < len(ids); i++ {
		ids[i] = worldsim.NewID()

		isUnique := true

		for j := 0; j < i; j++ {
			if ids[i] == ids[j] {
				isUnique = false

				break
			}
		}

		assert.True(t, isUnique)
	}
}
