package main

import (
    "fmt"
    "encoding/json"
)

//declaring a struct
type Human struct {
	//defining struct variables
	Name string
	Age  int
	Address string
}

//main function
func main() {
	
	//defining struct instance
	human1 := Human{"Ankit", 23, "New Delhi"}

	//encoding human1 struct
	//into json format
	human_enc, err := json.Marshal(human1)

	if err != nil {
		//if error is not nil, print error
		fmt.Println(err)
	}

	//as human_enc is in a byte array
	//format, it needs to be converted into a string
	fmt.Println(string(human_enc))

	//converting slices from golang into json format
	//defining an array of struct instance

	human2 := []Human {
		{Name : "Ashish", Age : 23, Address : "New Delhi"},
		{Name : "Rahul", Age : 24, Address : "Delhi"},
		{Name : "Pawan", Age : 22, Address : "Noida"},
	}
	//encoding into JSON format
	human2_enc, err := json.Marshal(human2)

	if err != nil {
		fmt.Println(err)
	}
	//printing encoded array
	fmt.Println()
	fmt.Println(string(human2_enc))
}