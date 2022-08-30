package entity

type inMemoryRepository struct {
	entities []*Entity
}

// Add entries to the manager.
func (a *inMemoryRepository) Add(entities ...*Entity) {
	a.entities = append(a.entities, entities...)
}

// Entities returns all the entities.
func (a *inMemoryRepository) Entities() (entities []*Entity) {
	return a.entities
}

// FilterByMask returns the mapped entities, which Components mask matched.
func (a *inMemoryRepository) FilterByMask(mask uint64) (entities []*Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*Entity, len(a.entities))
	index := 0
	for _, e := range a.entities {
		// Use the pre-calculated Components maskSlice.
		observed := e.Mask()
		// Add the entity to the filter list, if all Components are found.
		if observed&mask == mask {
			// Direct access
			entities[index] = e
			index++
		}
	}
	// Return only the needed slice.
	return entities[:index]
}

// FilterByNames returns the mapped entities, which Components names matched.
func (a *inMemoryRepository) FilterByNames(names ...string) (entities []*Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*Entity, len(a.entities))
	index := 0
	for _, e := range a.entities {
		// Each component should match
		matched := 0
		for _, name := range names {
			for _, c := range e.Components {
				switch v := c.(type) {
				case ComponentWithName:
					if v.Name() == name {
						matched++
					}
				}
			}
		}
		// Add the entity to the filter list, if all Components are found.
		if matched == len(names) {
			// Direct access
			entities[index] = e
			index++
		}
	}
	// Return only the needed slice.
	return entities[:index]
}

// Get a specific entity by Id.
func (a *inMemoryRepository) Get(id string) (entity *Entity) {
	for _, e := range a.entities {
		if e.ID() == id {
			return e
		}
	}
	return
}

// Remove a specific entity.
func (a *inMemoryRepository) Remove(entity *Entity) {
	for i, e := range a.entities {
		if e.Id == entity.Id {
			copy(a.entities[i:], a.entities[i+1:])
			a.entities[len(a.entities)-1] = nil
			a.entities = a.entities[:len(a.entities)-1]
			break
		}
	}
}

// NewInMemoryRepository creates a new inMemoryRepository and returns its address.
func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		entities: []*Entity{},
	}
}

// DefaultRepository is initialized by NewInMemoryRepository().
var DefaultRepository = NewInMemoryRepository()
