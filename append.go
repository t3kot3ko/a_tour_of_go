package main
import "fmt"

func main(){
	var a []int
	a = append(a, 100)
	fmt.Println(a)

	a = append(a, 200)
	fmt.Println(a)
}
