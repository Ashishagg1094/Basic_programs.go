//Go program to illustrate how to close a channel using for
//range loop and close function
package main

import "fmt"

//function
func myfun(mychnl chan string) {

    for v:=0; v<4; v++ {
        mychnl <- "Ashish"
    }
    close(mychnl)
}
//Main function
func main() {

    //Creating a channel 
    c := make(chan string)

    //calling Goroutine
    go myfun(c)

    for {
        res, ok := <-c
        if ok == false {
            fmt.Println("Channel Close ", ok)
            break
        }
        fmt.Println("Channel Open ", res, ok)
    }
}