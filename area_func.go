package main
import "fmt"

func main() {
	rect := area(2, 4)
	fmt.Println(rect)
}
	func area(l, b int) int {
		return l*b
	}