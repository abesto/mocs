package ecs

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMkAttributeWithUnknownType(t *testing.T) {
	type badType struct{}
	a, err := MkAttribute("badtype", reflect.TypeOf(badType{}).Name())
	assert.Nil(t, a)
	assert.EqualError(t, err, "Don't know how to create attribute for type badType")
}

func TestSetAttribute_InvalidValue(t *testing.T) {
	a, _ := MkAttribute("string-attribute", "string")
	value := 42
	err := a.Set(value)
	assert.Equal(t, "", a.Value())
	assert.EqualError(t, err, "Cannot assign value 42 to attribute string-attribute: int cannot be assigned to string")
}

func TestMkAttribute_string(t *testing.T) {
	a, err := MkAttribute("string-attribute", "string")
	assert.Nil(t, err)
	assert.Equal(t, "string-attribute", a.Name())
	assert.Equal(t, reflect.String, a.Type().Kind())
	assert.Equal(t, "", a.Value())
}

func TestSetAttribute_string(t *testing.T) {
	a, _ := MkAttribute("string-attribute", "string")
	value := "loremest of ipsums"
	assert.Nil(t, a.Set(value))
	assert.Equal(t, value, a.Value())
}

func TestSetAttribute_Entity(t *testing.T) {
	a, _ := MkAttribute("entity-attribute", "Entity")
	target := Entity(2131321)
	err := a.Set(target)
	assert.Nil(t, err)
	assert.Equal(t, target, a.Value())
}

func TestSetAttribute_uint64(t *testing.T) {
	a, _ := MkAttribute("int-attribute", "int")
	const target int64 = 919191
	err := a.Set(target)
	assert.Nil(t, err)
	assert.Equal(t, target, a.Value())
}
