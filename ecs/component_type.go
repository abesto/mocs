package ecs

import (
	"fmt"
)

// ComponentType describes a named set of attributes. The components of entities consist of instances of
// these component types. The value of attributes on a ComponentType describe the default value of those
// attributes, when a new component is instantiated on an entity.
type ComponentType struct {
	name       string
	attributes map[string]*Attribute
}

// MkComponentType creates a new component with no attributes.
func MkComponentType(name string) *ComponentType {
	return &ComponentType{name, map[string]*Attribute{}}
}

// CreateAttribute creates a new Attribute on this ComponentType.
func (t *ComponentType) CreateAttribute(name string, typeName string) error {
	if attr, exists := t.attributes[name]; exists {
		return fmt.Errorf("Component type %s already has an attribute called %s (type: %s; value: %#v)",
			t.name, name, attr.Type().Name(), attr.Value())
	}
	attr, err := MkAttribute(name, typeName)
	if err != nil {
		return err
	}
	t.attributes[name] = attr
	return nil
}

// SetDefault sets the default value of the named Attribute on this ComponentType.
func (t *ComponentType) SetDefault(name string, value interface{}) error {
	attr, exists := t.attributes[name]
	if exists {
		return attr.Set(value)
	}
	attributeNames := make([]string, 0, len(t.attributes))
	for attributeName := range t.attributes {
		attributeNames = append(attributeNames, attributeName)
	}
	return fmt.Errorf("Component type %s has no attribute called %s. Attributes it does have: %s",
		t.name, name, attributeNames)
}

// GetDefault returns the current default value of the named Attribute on this ComponentType.
func (t *ComponentType) GetDefault(name string) (interface{}, error) {
	attr, exists := t.attributes[name]
	if exists {
		return attr.Value(), nil
	}
	attributeNames := make([]string, 0, len(t.attributes))
	for attributeName := range t.attributes {
		attributeNames = append(attributeNames, attributeName)
	}
	return nil, fmt.Errorf("Component type %s has no attribute called %s. Attributes it does have: %s",
		t.name, name, attributeNames)
}

// RemoveAttribute removes a named Attribute from this ComponentType.
func (t *ComponentType) RemoveAttribute(name string) error {
	_, exists := t.attributes[name]
	if exists {
		delete(t.attributes, name)
		return nil
	}
	attributeNames := make([]string, 0, len(t.attributes))
	for attributeName := range t.attributes {
		attributeNames = append(attributeNames, attributeName)
	}
	return fmt.Errorf("Component type %s has no attribute called %s. Attributes it does have: %s",
		t.name, name, attributeNames)
}
