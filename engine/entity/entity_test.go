package entity_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/engine/entity"
	"testing"
)

func TestEntity_NewEntity_Should_Create_A_Correct_Mask(t *testing.T) {
	e := entity.NewEntity("e", []entity.Component{
		&mockComponent{name: "position", mask: 1},
	})
	assert.That("mask should be 1", t, e.Mask(), 1)
}

func TestEntity_Add_Should_Work_With_Multiple_Components(t *testing.T) {
	e := entity.NewEntity("e", []entity.Component{
		&mockComponent{name: "position", mask: 1},
	})
	e.Add(&mockComponent{name: "velocity", mask: 2})
	assert.That("mask should be 3", t, e.Mask(), 3)
}

func TestEntity_Remove_Should_Work_With_Multiple_Components(t *testing.T) {
	e := entity.NewEntity("e", []entity.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
		&mockComponent{name: "velocity", mask: 4},
	})
	e.Remove(4)
	assert.That("mask should be 3", t, e.Mask(), 3)
}

type mockComponent struct {
	mask uint64
	name string
}

func (c *mockComponent) Mask() uint64 { return c.mask }
func (c *mockComponent) Name() string { return c.name }
