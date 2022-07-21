package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Place struct {
	Name    string
	Country string
}

type Car struct {
	Make  string
	Model string
}

var container = CreateContainer()

func get[T interface{}]() T {
	return Get[T](container)
}
func single[T interface{}](b DependencyBuilder) {
	Single[T](container, b)
}
func factory[T interface{}](b DependencyBuilder) {
	Factory[T](container, b)
}

type ComplexPerson struct {
	Person
	Car  Car
	Home Place
}

func main() {
	factory[ComplexPerson](func() interface{} {
		return ComplexPerson{
			get[Person](),
			get[Car](),
			get[Place](),
		}
	})
	factory[Person](func() interface{} {
		return Person{"Jason", 24}
	})
	factory[Place](func() interface{} {
		return Place{"New York", "USA"}
	})
	factory[Car](func() interface{} {
		return Car{"Toyota", "Hilux"}
	})
	single[*Car](func() interface{} {
		return &Car{"Toyota", "Yaris"}
	})

	cp1 := get[ComplexPerson]()
	cp2 := get[ComplexPerson]()

	fmt.Printf("cp1 => %#v\n", cp1)
	fmt.Printf("cp2 => %#v\n", cp2)
	fmt.Println()
	fmt.Printf("cp1:Address => %p\n", &cp1)
	fmt.Printf("cp2:Address => %p\n", &cp2)
}
