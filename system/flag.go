package system

import "sync"

type Flag struct {
	value bool
	mutex sync.Mutex
}

func NewFlag(val bool) *Flag {
	return &Flag{
		value: val,
	}
}

func (f *Flag) Get() bool {
	f.mutex.Lock()

	defer f.mutex.Unlock()

	return f.value
}

func (f *Flag) Set(val bool) {
	f.mutex.Lock()

	defer f.mutex.Unlock()

	f.value = val
}
