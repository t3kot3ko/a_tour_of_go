package main
import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m["Answer"])

	delete(m, "Answerx")
	fmt.Println(m["Answer"])

	v, ok := m["Answer"]
	fmt.Println(v, ok)
}
