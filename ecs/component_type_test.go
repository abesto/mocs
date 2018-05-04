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

func TestComponentTypeDefault(t *testing.T) {
	c := MkComponentType("default-test")

	def0, err0 := c.GetDefault("test-attr")
	assert.Nil(t, def0)
	assert.EqualError(t, err0,
		"Component type default-test has no attribute called test-attr. Attributes it does have: []")

	c.CreateAttribute("test-attr", "bool")
	def1, err1 := c.GetDefault("test-attr")
	assert.Nil(t, err1)
	assert.False(t, def1.(bool))

	c.SetDefault("test-attr", true)
	def2, err2 := c.GetDefault("test-attr")
	assert.Nil(t, err2)
	assert.True(t, def2.(bool))
}

func TestComponentTypeRemoveAttribute(t *testing.T) {
	c := MkComponentType("remove-test")
	assert.Nil(t, c.CreateAttribute("dummy-attr", "string"))
	assert.EqualError(t, c.RemoveAttribute("test-attr"),
		"Component type remove-test has no attribute called test-attr. Attributes it does have: [dummy-attr]")

	c.CreateAttribute("test-attr", "bool")
	def0, err0 := c.GetDefault("test-attr")
	assert.Nil(t, err0)
	assert.False(t, def0.(bool))

	assert.Nil(t, c.RemoveAttribute("test-attr"))
	def1, err1 := c.GetDefault("test-attr")
	assert.Nil(t, def1)
	assert.EqualError(t, err1,
		"Component type remove-test has no attribute called test-attr. Attributes it does have: [dummy-attr]")

	assert.EqualError(t, c.RemoveAttribute("test-attr"),
		"Component type remove-test has no attribute called test-attr. Attributes it does have: [dummy-attr]")
}
