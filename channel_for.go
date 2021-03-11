package main

import "fmt"

//Main function
func main() {

    //Creating a channel
    //Using make() function
    mychnl := make(chan string)

    //Anonymouss Goroutine 
    go func() {
        mychnl <- "A"
        mychnl <- "S"
        mychnl <- "H"
        mychnl <- "I"
        mychnl <- "S"
        mychnl <- "H"
        close(mychnl)
    } ()

    //Using for loop
    for res := range mychnl {
        fmt.Println(res)
    }
    
}