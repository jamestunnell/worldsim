package system_test

import (
	"sync"
	"testing"

	"github.com/jamestunnell/worldsim/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTwoSystems(t *testing.T) {
	intChan := make(chan int, 20)
	s1 := system.New("s1", func() {
		intChan <- 1
	})
	s2 := system.New("s2", func() {
		intChan <- 2
	})

	require.NoError(t, system.StartAll(s1, s2))

	defer system.StopAll(s1, s2)

	var wg sync.WaitGroup

	wg.Add(1)

	s1.Exec(func() { wg.Done() })

	wg.Wait()

	wg.Add(1)

	s2.Exec(func() { wg.Done() })

	wg.Wait()

	intVals := []int{}
	anyLeft := true
	for anyLeft {
		select {
		case intVal := <-intChan:
			intVals = append(intVals, intVal)
		default:
			anyLeft = false
		}
	}

	require.Len(t, intVals, 2)

	assert.Equal(t, 1, intVals[0])
	assert.Equal(t, 2, intVals[1])
}
