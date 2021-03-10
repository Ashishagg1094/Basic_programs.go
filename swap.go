package main
import "fmt"

func main() {
	a,b := 2,4
	c := a+b
	a = c-a
	b = c-b
	fmt.Println(a)
	fmt.Println(b)
}