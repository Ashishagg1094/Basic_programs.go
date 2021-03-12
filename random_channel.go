package main

import (
    "fmt"
	"time"
    "math/rand"
)

func myfunc(ch chan int) {
	time.Sleep(time.Second)
	for j:=0; j<5; j++ {
	fmt.Println(<-ch)
	}
	
}
func main() {
    
    //creating a channel
    ch := make(chan int)
    
	go myfunc(ch)
	for i:=0; i<5; i++ {
	ch <- rand.Intn(10)
	}
}