package engine

import "github.com/andygeiss/utils/engine/entity"

// Plugin is a function which handles a specific kind of functionality
// by using an Entity Repository to gain access to the entities.
type Plugin func(er entity.Repository)
