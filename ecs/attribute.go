package ecs

import (
	"fmt"
	"reflect"
)

// Attribute describes one typed attribute of a component
type Attribute struct {
	name  string
	ttype reflect.Type
	value interface{}
}

const cInt64 int64 = 0

var (
	typeEntity = reflect.TypeOf(Entity(0))
	typeInt64  = reflect.TypeOf(cInt64)
)

func tryCast(value interface{}, targetType reflect.Type) interface{} {
	fromType := reflect.TypeOf(value)
	if fromType == typeInt64 && targetType == typeEntity {
		return Entity(value.(uint64))
	}
	return value
}

func (a *Attribute) typeCheck(value interface{}) (interface{}, error) {
	expectedType := a.ttype
	castValue := tryCast(value, expectedType)
	actualType := reflect.TypeOf(castValue)
	if !actualType.AssignableTo(expectedType) {
		return nil, fmt.Errorf("%s cannot be assigned to %s", actualType, expectedType)
	}
	return castValue, nil
}

// Set the value of this attribute, with run-time type checking
func (a *Attribute) Set(value interface{}) error {
	castValue, err := a.typeCheck(value)
	if err != nil {
		return fmt.Errorf("Cannot assign value %#v to attribute %s: %s",
			value, a.name, err,
		)
	}
	a.value = castValue
	return nil
}

// Name returns the name of this attribute
func (a *Attribute) Name() string {
	return a.name
}

// Type returns the type of this attribute
func (a *Attribute) Type() reflect.Type {
	return a.ttype
}

// Value returns the current value of this attribute
func (a *Attribute) Value() interface{} {
	return a.value
}

// MkAttribute creates a new Attribute with a given name and run-time type
func MkAttribute(name string, typeName string) (*Attribute, error) {
	var value interface{}
	var t reflect.Type
	if typeName == "string" {
		value = ""
	} else if typeName == "int" {
		const n int64 = 0
		value = n
	} else if typeName == "bool" {
		value = false
	} else if typeName == "Entity" || typeName == "ecs.Entity" {
		value = nil
		t = typeEntity
	} else {
		return nil, fmt.Errorf("Don't know how to create attribute for type %s", typeName)
	}
	if t == nil {
		t = reflect.TypeOf(value)
	}
	return &Attribute{name, t, value}, nil
}
