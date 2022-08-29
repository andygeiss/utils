package engine

// Engine describes a solution that is divided into systems to implement different aspects of a complex problem.
type Engine interface {
	Setup()
	Start() (stopCh chan bool)
	State() (state int)
	Teardown()
	WithSystems(s ...System) Engine
}

const (
	StateEngineStopped = iota
	StateEngineRunning
)
