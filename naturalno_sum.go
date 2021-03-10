package main
import "fmt"

func main() {
	sum := 0
	num:=10
	for i:=1; i<=num; i++ {
		sum+=i
	}
	fmt.Printf("The sum of first %v natural numbers is: %v", num, sum)
}