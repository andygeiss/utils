package engine

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// defaultEngine implements a straightforward implementation using channels.
type defaultEngine struct {
	sigTerm chan os.Signal
	state   int
	systems []System
	mutex   sync.Mutex
}

func (a *defaultEngine) Setup() (stopCh chan bool) {
	// set up a goroutine which waits for SIGTERM to close the program.
	a.sigTerm = make(chan os.Signal, 2)
	signal.Notify(a.sigTerm, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-a.sigTerm
		a.Teardown()
		os.Exit(0)
	}()
	// set up a goroutine which waits for a stop.
	ch := make(chan bool)
	go func() {
		<-ch
		a.state = StateEngineStopped
	}()
	// first set up the systems.
	for _, sys := range a.systems {
		sys.Setup()
	}
	// set the state
	a.state = StateEngineRunning
	// set up a goroutine for a loop to process the systems.
	go func() {
		for _, sys := range a.systems {
			sys.Process(ch)
		}
	}()
	return ch
}

func (a *defaultEngine) State() (state int) {
	return a.state
}

func (a *defaultEngine) Teardown() {
	for _, sys := range a.systems {
		sys.Teardown()
	}
	a.state = StateEngineStopped
}

func (a *defaultEngine) WithSystems(s ...System) Engine {
	a.systems = append(a.systems, s...)
	return a
}

func NewDefaultEngine() Engine {
	return &defaultEngine{
		state: StateEngineStopped,
	}
}

var DefaultEngine = NewDefaultEngine()
