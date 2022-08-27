package entity_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/engine/entity"
	"testing"
)

func TestEntityManager_Entities_Should_Have_No_Entity_At_Start(t *testing.T) {
	m := entity.NewInMemoryRepository()
	assert.That("manager should have no entity at start", t, len(m.Entities()), 0)
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Adding_One_Entity(t *testing.T) {
	m := entity.NewInMemoryRepository()
	m.Add(&entity.Entity{})
	assert.That("manager should have one entity", t, len(m.Entities()), 1)
}

func TestEntityManager_Entities_Should_Have_Two_Entities_After_Adding_Two_Entities(t *testing.T) {
	m := entity.NewInMemoryRepository()
	m.Add(entity.NewEntity())
	m.Add(entity.NewEntity())
	assert.That("manager should have two entity", t, len(m.Entities()), 2)
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Removing_One_Of_Two_Entities(t *testing.T) {
	m := entity.NewInMemoryRepository()
	e1 := entity.NewEntity().WithID("e1")
	e2 := entity.NewEntity().WithID("e2")
	m.Add(e1)
	m.Add(e2)
	m.Remove(e2)
	assert.That("manager should have one entity after removing one out of two", t, len(m.Entities()), 1)
	assert.That("remaining entity should have Id e1", t, m.Entities()[0].ID(), "e1")
}

func TestEntityManager_FilterByMask_Should_Return_No_Entity_Out_Of_One(t *testing.T) {
	em := entity.NewInMemoryRepository()
	e := entity.NewEntity().
		WithID("e1").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
		)
	em.Add(e)
	filtered := em.FilterByMask(2)
	assert.That("filter should return no entity", t, len(filtered), 0)
}

func TestEntityManager_FilterByMask_Should_Return_One_Entity_Out_Of_One(t *testing.T) {
	em := entity.NewInMemoryRepository()
	e := entity.NewEntity().
		WithID("e1").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
		)
	em.Add(e)
	filtered := em.FilterByMask(1)
	assert.That("filter should return one entity", t, len(filtered), 1)
}

func TestEntityManager_FilterByMask_Should_Return_One_Entity_Out_Of_Two(t *testing.T) {
	em := entity.NewInMemoryRepository()
	e1 := entity.NewEntity().
		WithID("e1").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
		)
	e2 := entity.NewEntity().
		WithID("e2").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	em.Add(e1, e2)
	filtered := em.FilterByMask(2)
	assert.That("filter should return one entity", t, len(filtered), 1)
	assert.That("entity should be e2", t, filtered[0], e2)
}

func TestEntityManager_FilterByMask_Should_Return_Two_Entities_Out_Of_Three(t *testing.T) {
	em := entity.NewInMemoryRepository()
	e1 := entity.NewEntity().
		WithID("e1").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
		)
	e2 := entity.NewEntity().
		WithID("e2").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	e3 := entity.NewEntity().
		WithID("e3").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	em.Add(e1, e2, e3)
	filtered := em.FilterByMask(2)
	assert.That("filter should return one entity", t, len(filtered), 2)
	assert.That("entity should be e2", t, filtered[0], e2)
	assert.That("entity should be e3", t, filtered[1], e3)
}

func TestEntityManager_FilterByMask_Should_Return_Three_Entities_Out_Of_Three(t *testing.T) {
	em := entity.NewInMemoryRepository()
	e1 := entity.NewEntity().
		WithID("e1").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	e2 := entity.NewEntity().
		WithID("e2").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	e3 := entity.NewEntity().
		WithID("e3").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	em.Add(e1, e2, e3)
	filtered := em.FilterByMask(1 | 2)
	assert.That("filter should return one entity", t, len(filtered), 3)
	assert.That("entity should be e1", t, filtered[0], e1)
	assert.That("entity should be e2", t, filtered[1], e2)
	assert.That("entity should be e3", t, filtered[2], e3)
}

func TestEntityManager_FilterByNames_Should_Return_Three_Entities_Out_Of_Three(t *testing.T) {
	em := entity.NewInMemoryRepository()
	e1 := entity.NewEntity().
		WithID("e1").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	e2 := entity.NewEntity().
		WithID("e2").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	e3 := entity.NewEntity().
		WithID("e3").
		WithComponents(
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
		)
	em.Add(e1, e2, e3)
	filtered := em.FilterByNames("position", "size")
	assert.That("filter should return one entity", t, len(filtered), 3)
	assert.That("entity should be e1", t, filtered[0], e1)
	assert.That("entity should be e2", t, filtered[1], e2)
	assert.That("entity should be e3", t, filtered[2], e3)
}
