package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go is running on")

	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
		fallthrough
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other")
	}
}


