package main

import (
	"fmt"
	"reflect"
)

type DependencyBuilder func() interface{}

type Container struct {
	singleBuilders  map[string]DependencyBuilder
	factoryBuilders map[string]DependencyBuilder
	cache           map[string]any
}

func CreateContainer() *Container {
	return &Container{
		make(map[string]DependencyBuilder),
		make(map[string]DependencyBuilder),
		make(map[string]interface{}),
	}
}

func Factory[T any](c *Container, b DependencyBuilder) {
	var temp T
	c.factoryBuilders[reflect.TypeOf(temp).String()] = b
}

func Single[T any](c *Container, b DependencyBuilder) {
	var dep T
	c.singleBuilders[reflect.TypeOf(dep).String()] = b
}

func Get[T any](c *Container) T {
	var temp T
	typeStr := reflect.TypeOf(temp).String()

	if dep, ok := c.cache[typeStr]; ok {
		return dep.(T)
	}

	if builder, ok := c.singleBuilders[typeStr]; ok {
		temp = builder().(T)
		c.cache[typeStr] = temp
		return temp
	}

	if builder, ok := c.factoryBuilders[typeStr]; ok {
		return builder().(T)
	}

	panic(fmt.Sprintf("dependency of type  '%s' not found in container", typeStr))
}
