package ecs

import (
	"fmt"
)

type ComponentType struct {
	name       string
	attributes map[string]*Attribute
}

func MkComponentType(name string) *ComponentType {
	return &ComponentType{name, map[string]*Attribute{}}
}

func (t *ComponentType) CreateAttribute(name string, typeName string) error {
	if attr, exists := t.attributes[name]; exists {
		return fmt.Errorf("Component type %s already has an attribute called %s (type: %s; value: %#v)",
			t.name, name, attr.Type().Name(), attr.Value())
	}
	if attr, err := MkAttribute(name, typeName); err != nil {
		return err
	} else {
		t.attributes[name] = attr
		return nil
	}
}

func (t *ComponentType) SetAttribute(name string, value interface{}) error {
	if attr, exists := t.attributes[name]; exists {
		return attr.Set(value)
	} else {
		attributeNames := make([]string, 0, len(t.attributes))
		for attributeName, _ := range t.attributes {
			attributeNames = append(attributeNames, attributeName)
		}
		return fmt.Errorf("Component type %s has no attribute called %s. Attributes it does have: %s",
			t.name, name, attributeNames)
	}
}
