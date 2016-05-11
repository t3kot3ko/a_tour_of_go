package main
import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	var t int

	return func() int {
		t = y
		y = x + y
		x = t
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
