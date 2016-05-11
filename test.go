package main

import "fmt"

type Foo struct {
	X, Y int
}

func (f *Foo) Double() {
	f.X = f.X * 2
	f.Y = f.Y * 2
}

func (f Foo) String() string {
	return fmt.Sprintf("(%d, %d)", f.X, f.Y)
}

func main() {
	var f Foo = Foo{10, 20}
	fmt.Println(f.X, f.Y)
	f.Double()
	fmt.Println(f.X, f.Y)

	fmt.Println(f)
}
