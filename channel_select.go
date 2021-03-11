//Go program to illustrate the concept of select statement
package main

import(
    "fmt"
    "time"
)

//function1
func portal1(channel1 chan string) {
    time.Sleep(9*time.Second)
    channel1 <- "Welcome to channel1"
}

//function2
func portal2(channel2 chan string) {
    time.Sleep(3*time.Second)
    channel2 <- "Welcome to channel2"
}

//main function
func main() {

    //Creating channels
    R1 := make(chan string)
    R2 := make(chan string)

    //Calling function1 and function2 in goroutine
    go portal1(R1)
    go portal2(R2)

    select {

        //case 1 for portal1
        case op1 := <- R1:
        fmt.Println(op1)

        //case 2 for portal2
        case op2 := <- R2:
        fmt.Println(op2)
    }
    
}