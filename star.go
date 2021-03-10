package main
import "fmt"

func main() {
	for i:=1; i<=5 ;i++ {
		for j:=i; j<=5; j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}