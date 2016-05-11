package main

import "fmt"

func fibonacchi(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)

	c <- 199
}

func main() {
	c := make(chan int, 10)
	go fibonacchi(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
