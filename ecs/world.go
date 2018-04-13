package ecs

type World struct {
	Entities       map[Entity]bool
	ComponentTypes map[uint64]*ComponentType
}
