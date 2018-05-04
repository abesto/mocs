package ecs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMkWorld(t *testing.T) {
	w, err := MkWorld(42)
	assert.Nil(t, err)
	assert.Empty(t, w.entities)
	assert.Empty(t, w.componentTypes)
}

func TestWorldCreateEntity(t *testing.T) {
	w, err := MkWorld(1)
	assert.Nil(t, err)

	w.entityLimit = 1

	entity, err := w.CreateEntity()
	assert.Nil(t, err)
	assert.True(t, w.entities[entity])

	_, err = w.CreateEntity()
	assert.EqualError(t, err, "World already has 1 entities, unable to create more")
}

func TestWorldDeleteEntity(t *testing.T) {
	w, err := MkWorld(2)
	assert.Nil(t, err)

	entity0, err := w.CreateEntity()
	assert.Nil(t, err)
	assert.True(t, w.entities[entity0])
	entity1, err := w.CreateEntity()
	assert.Nil(t, err)
	assert.True(t, w.entities[entity1])

	assert.Nil(t, w.DeleteEntity(entity0))
	assert.False(t, w.entities[entity0])
	assert.True(t, w.entities[entity1])

	// TODO test that components are deleted
}
