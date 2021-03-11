package main
import "fmt"

func main() {

    //Creating a channel using var keyword
    var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel)
    fmt.Printf("Type of the channel: %T", mychannel)

    //Creating a channel using make() function
    mychannel1 := make(chan int)
    fmt.Println("\nValue of the channel1: ", mychannel1)
    fmt.Println("Type of the channel1 %T: ", mychannel1) 
}