package main

import (
	"fmt"
	"reflect"
)

type DependencyBuilder func() interface{}

type Container struct {
	singleBuilders  map[string]DependencyBuilder
	factoryBuilders map[string]DependencyBuilder
	cache           map[string]interface{}
}

func CreateContainer() *Container {
	return &Container{
		make(map[string]DependencyBuilder),
		make(map[string]DependencyBuilder),
		make(map[string]interface{}),
	}
}

func Factory[T interface{}](c *Container, b DependencyBuilder) {
	var temp T
	c.factoryBuilders[reflect.TypeOf(temp).String()] = b
}

func Single[T interface{}](c *Container, b DependencyBuilder) {
	var dep T
	c.singleBuilders[reflect.TypeOf(dep).String()] = b
}

func Instance[T interface{}](c *Container, dep T) {
	c.cache[reflect.TypeOf(dep).String()] = dep
}

func Get[T interface{}](c *Container) T {
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
