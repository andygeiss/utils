package engine_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/engine"
	"testing"
	"time"
)

func TestDefaultEngine_Setup(t *testing.T) {
	var valueSetup, stateSetup int
	var valueProcess, stateProcess int
	sys := &mockupSystem{}
	eng := engine.NewDefaultEngine().WithSystems(sys)
	eng.Setup()
	valueSetup = sys.Value
	stateSetup = eng.State()
	eng.Start()
	valueProcess = sys.Value
	stateProcess = eng.State()
	assert.That("state should be StateEngineStopped after calling Setup", t, stateSetup, engine.StateEngineStopped)
	assert.That("state should be StateEngineStopped after calling Process", t, stateProcess, engine.StateEngineStopped)
	assert.That("valueSetup should be 0", t, valueSetup, 0)
	assert.That("valueProcess should be 1", t, valueProcess, 1)
}

func TestDefaultEngine_Teardown(t *testing.T) {
	sys := &mockupSystem{}
	eng := engine.NewDefaultEngine().WithSystems(sys)
	eng.Setup()
	eng.Teardown()
	state := eng.State()
	assert.That("state should be StateEngineStopped after teardown", t, state, engine.StateEngineStopped)
	assert.That("value should be 2", t, sys.Value, 2)
}

type mockupSystem struct {
	Value int
}

func (a *mockupSystem) Error() (err error) { return nil }
func (a *mockupSystem) Process(stopCh chan bool) {
	a.Value = 1
	stopCh <- true
	time.Sleep(time.Microsecond * 1)
}
func (a *mockupSystem) Setup() {
	a.Value = 0
}
func (a *mockupSystem) Teardown() {
	a.Value = 2
}
