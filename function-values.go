package main
import (
	"fmt"
	"math"
)

func compute(f func(float64,float64) float64) float64 {
	return f(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}

	var f func(float64, float64) float64 = func(x, y float64) float64{
		return x + y
	}

	fmt.Println( compute(hypot) )
	fmt.Println( compute(f) )
}
