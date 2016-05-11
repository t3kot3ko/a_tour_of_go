package main

import "fmt"

type Vertex struct{
	X int
	Y int
}

func main() {
	v := Vertex{100, 200}
	p := &v
	p.X = 300
	fmt.Println(v)
}
