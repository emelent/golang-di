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

type Shoe struct {
	Brand string
	Size  int
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
func instance[T interface{}](dep T) {
	Instance[T](container, dep)
}

func main() {
	factory[Person](func() interface{} {
		return Person{"Jason", 24}
	})
	factory[Place](func() interface{} {
		return Place{"New York", "USA"}
	})
	single[*Car](func() interface{} {
		return &Car{"Toyota", "Yaris"}
	})
	instance[*Shoe](&Shoe{"Nike", 10})

	p1 := get[Person]()
	p2 := get[Person]()

	c1 := get[*Car]()
	c2 := get[*Car]()

	s1 := get[*Shoe]()
	s2 := get[*Shoe]()

	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)
	fmt.Printf("p1 == p2 => %#v\n", p1 == p2)
	fmt.Printf("&p1 == &p2 => %#v\n", &p1 == &p2)
	fmt.Println()

	fmt.Printf("%#v\n", c1)
	fmt.Printf("%#v\n", c2)
	fmt.Printf("c1 == c2 => %#v\n", *c1 == *c2)
	fmt.Printf("&c1 == &c2 => %#v\n", c1 == c2)
	fmt.Println()

	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)
	fmt.Printf("s1 == s2 => %#v\n", *s1 == *s2)
	fmt.Printf("&s1 == &s2 => %#v\n", s1 == s2)

}
