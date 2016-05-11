package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main(){
	// var m map[string] Vertex = make(map[string]Vertex)
	m := make(map[string]Vertex)

	m["tokyo"] = Vertex{123.0, 234.0}
	fmt.Println(m)
}

