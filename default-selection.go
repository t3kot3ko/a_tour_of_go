package main

import "fmt"
import "time"

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case x := <-tick:
			fmt.Println("tick")
			fmt.Println(x)
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("       .")
			time.Sleep(60 * time.Millisecond)
		}
	}
}
