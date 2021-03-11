package main
import "fmt"

	type Animal struct {
		Name string
		Origin string
	}
	type Bird struct {
		Animal
		SpeedKPH float32
		Canfly bool
	}
	func main() {
		b := Bird{}
		b.Name = "Emu"
		b.Origin = "Australia"
		b.SpeedKPH = 48
		b.Canfly = false
		fmt.Println(b)
}