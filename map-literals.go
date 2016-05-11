package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex {
		"Tokyo": Vertex{100, 200},
		"US": Vertex{300, 400},
	}

	fmt.Println(m)
}
