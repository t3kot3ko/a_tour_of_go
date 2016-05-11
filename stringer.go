package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Foo", 100}
	b := Person{"Bar", 200}

	fmt.Println(a, b)
	fmt.Println(a.String())
}
