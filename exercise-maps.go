package main

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	fields := strings.Fields(s)
	for _, f := range fields {
		result[f]++
	}
	return result
}

func main() {
	s := "foo bar baz bar"
	fmt.Println( WordCount(s) )
}
