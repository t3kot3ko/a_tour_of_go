package main

import (
	"fmt"
	"time"
)

func main(){
	today := time.Now().Weekday()
	fmt.Println(today)

	switch time.Saturday {
	case today + 1:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Other")

	}
	
}
