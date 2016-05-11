package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{100, 200}
	v.X = 3
	fmt.Println(v.X)
}
