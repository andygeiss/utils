package entity

// Component contains only the data (no behavior).
// The Name() method must be implemented, because the inMemoryRepository
// uses it to filter the entities by component name.
type Component interface {
	Mask() uint64
}

// ComponentWithName is used by FilterByNames to enable more than 64 Components (if needed).
type ComponentWithName interface {
	Component
	Name() string
}

// Entity is simply a composition of one or more Components with an Id.
type Entity struct {
	Components []Component `json:"components"`
	Id         string      `json:"id"`
	Masked     uint64      `json:"masked"`
}

// Add a component.
func (a *Entity) Add(cn ...Component) {
	for _, c := range cn {
		if a.Masked&c.Mask() == c.Mask() {
			continue
		}
		a.Components = append(a.Components, c)
		a.Masked = maskSlice(a.Components)
	}
}

// Get a component by its bitmask.
func (a *Entity) Get(mask uint64) Component {
	for _, c := range a.Components {
		if c.Mask() == mask {
			return c
		}
	}
	return nil
}

// ID ...
func (a *Entity) ID() string {
	return a.Id
}

// Mask returns a pre-calculated maskSlice to identify the Components.
func (a *Entity) Mask() uint64 {
	return a.Masked
}

// Remove a component by using its maskSlice.
func (a *Entity) Remove(mask uint64) {
	modified := false
	for i, c := range a.Components {
		if c.Mask() == mask {
			copy(a.Components[i:], a.Components[i+1:])
			a.Components[len(a.Components)-1] = nil
			a.Components = a.Components[:len(a.Components)-1]
			modified = true
			break
		}
	}
	if modified {
		a.Masked = maskSlice(a.Components)
	}
}

func (a *Entity) WithComponents(components ...Component) *Entity {
	a.Add(components...)
	return a
}

func (a *Entity) WithID(id string) *Entity {
	a.Id = id
	return a
}

// NewEntity creates a new entity and pre-calculates the component maskSlice.
func NewEntity() *Entity {
	return &Entity{}
}

func maskSlice(components []Component) uint64 {
	mask := uint64(0)
	for _, c := range components {
		mask = mask | c.Mask()
	}
	return mask
}
