package main
import(
    "fmt"
    "time"
)
func main() {
    fmt.Println("Welcome!! to Main function")

    //Creating Anonymous Goroutine
    go func() {
        fmt.Println("Welcome!! to Ashish")
    }()

    time.Sleep(1*time.Second)
    fmt.Println("GoodBye!! to Main function")
}