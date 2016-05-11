package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	if today == time.Friday {
		fmt.Println("Friday")
	} else if today == time.Thursday {
		fmt.Println("Thursday")
	}
}
