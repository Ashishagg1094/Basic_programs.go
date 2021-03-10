package main
import "fmt"

func main() {
    fact := 1
	num := 5
	for i:=1; i<=num; i++ {
		fact*=i
	}
	fmt.Println(fact)
}