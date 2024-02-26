package system

type System interface {
	Start() error
	Stop()
	Running() bool
	Exec(onComplete func())
}
