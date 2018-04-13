package ecs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponentTypeAttributeDoubleCreate(t *testing.T) {
	c := MkComponentType("double-test")
	assert.Nil(t, c.CreateAttribute("test-attr", "string"))
	assert.EqualError(t,
		c.CreateAttribute("test-attr", "string"),
		"Component type double-test already has an attribute called test-attr (type: string; value: \"\")")
}
