package system

import (
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/worldsim/system"
)

type System struct {
	running *system.Flag
	stop    chan struct{}
	exec    chan func()
}

var errAlreadyRunning = errors.New("physics system is already running")

func New(name string, do func()) *System {
	return &System{
		running: system.NewFlag(false),
		stop:    make(chan struct{}),
		exec:    make(chan func()),
	}
}

func (sys *System) Start() error {
	if sys.running.Get() {
		return errAlreadyRunning
	}

	go sys.run()

	return nil
}

func (sys *System) Stop() {
	if !sys.running.Get() {
		return
	}

	sys.stop <- struct{}{}
}

func (sys *System) Running() bool {
	return sys.running.Get()
}

func (sys *System) Exec(onComplete func()) {
	sys.exec <- onComplete
}

func (sys *System) run() {
	log.Info().Msg("physics system started")

	select {
	case <-sys.stop:
		break
	case onComplete := <-sys.exec:
		sys.do()

		onComplete()
	}

	log.Info().Msg("physics system stopped")
}

func (sys *System) do() {

}
