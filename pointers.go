package main

import "fmt"

func main(){
	i, j := 12, 23
	p := &i
	fmt.Println(*p)
	fmt.Println(j)
}
