package ecs

import (
	"fmt"
	"math"

	"github.com/bwmarrin/snowflake"
)

// World is the root of the ECS system.
type World struct {
	entities       map[Entity]bool
	componentTypes map[string]*ComponentType
	snowflakeNode  *snowflake.Node
	entityLimit    uint64
}

// MkWorld creates a brand-new world.
func MkWorld(nodeID int64) (*World, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}
	return &World{
		entities:       map[Entity]bool{},
		componentTypes: map[string]*ComponentType{},
		snowflakeNode:  node,
		entityLimit:    math.MaxUint64,
	}, nil
}

func (w *World) isEntityIDFree(entity Entity) bool {
	_, exists := w.entities[entity]
	return !exists
}

func (w *World) getFreeEntityID() (Entity, error) {
	if uint64(len(w.entities)) >= w.entityLimit {
		return 0, fmt.Errorf("World already has %d entities, unable to create more", len(w.entities))
	}
	var entity Entity
	for entity = Entity(w.snowflakeNode.Generate()); !w.isEntityIDFree(entity); {
		entity = Entity(w.snowflakeNode.Generate())
	}
	return entity, nil
}

// CreateEntity creates a new Entity in the World.
func (w *World) CreateEntity() (Entity, error) {
	entity, err := w.getFreeEntityID()
	if err == nil {
		w.entities[entity] = true
	}
	return entity, err
}

// DeleteEntity deletes an Entity from the World, and cleans up all its Components.
func (w *World) DeleteEntity(entity Entity) error {
	if w.isEntityIDFree(entity) {
		return fmt.Errorf("No entity with ID %d exists", entity)
	}
	delete(w.entities, entity)
	return nil
}
