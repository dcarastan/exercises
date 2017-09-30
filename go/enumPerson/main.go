package main

import (
	"fmt"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func newPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func main() {
	words := []string{"foo", "bar", "zoo"}
	fmt.Println("Hello")
	for i, w := range words {
		fmt.Printf("%d: %s\n", i, w)
	}

	p := newPerson("Doru", 50)

	fmt.Printf("%v\n", p)
}
