package main
import "fmt"

func main() {
	s := []int{2, 3, 5, 6, 7}
	fmt.Println(s[1:3])
	fmt.Println(s[1:])
	fmt.Println(s[len(s) - 1])
	fmt.Println(s[10])

}
