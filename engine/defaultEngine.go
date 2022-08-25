package engine

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// defaultEngine implements a simple implementation using channels.
type defaultEngine struct {
	sigTerm chan os.Signal
	state   int
	systems []System
	mutex   sync.Mutex
}

// Setup initializes the engine and its subsystems.
func (a *defaultEngine) Setup() (stopCh chan bool) {
	// Set up a goroutine that waits for SIGTERM to terminate the program.
	a.setupSigTerm()
	// Set up a goroutine that waits for a stop.
	ch := a.setupStopCh()
	// First, set up the systems.
	for _, sys := range a.systems {
		sys.Setup()
	}
	// Set the initial state.
	a.state = StateEngineRunning
	// Set up a goroutine for a loop to process the systems.
	go func() {
		for a.state == StateEngineRunning {
			for _, sys := range a.systems {
				sys.Process(ch)
			}
		}
	}()
	return ch
}

// State returns the current state of the engine.
func (a *defaultEngine) State() (state int) {
	return a.state
}

// Teardown shuts down the engine and its subsystems.
func (a *defaultEngine) Teardown() {
	for _, sys := range a.systems {
		sys.Teardown()
	}
	a.state = StateEngineStopped
}

// WithSystems adds a specific number of systems to the engine.
func (a *defaultEngine) WithSystems(s ...System) Engine {
	a.systems = append(a.systems, s...)
	return a
}

func (a *defaultEngine) setupSigTerm() {
	a.sigTerm = make(chan os.Signal, 2)
	signal.Notify(a.sigTerm, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-a.sigTerm
		a.state = StateEngineStopped
	}()
}

func (a *defaultEngine) setupStopCh() chan bool {
	ch := make(chan bool)
	go func() {
		<-ch
		a.state = StateEngineStopped
	}()
	return ch
}

// NewDefaultEngine creates a new engine an returns its address.
func NewDefaultEngine() Engine {
	return &defaultEngine{
		state: StateEngineStopped,
	}
}

// DefaultEngine ...
var DefaultEngine = NewDefaultEngine()
