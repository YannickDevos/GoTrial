package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5} // create a slice of int with a length of 5 and a capacity of 5
	fmt.Println(mySlice)

	for _, i := range mySlice { //loop through myslice and print each value of the slice
		fmt.Println("for range loop:", i)
	}

	newslice := []int{10, 20, 30}          //create a new slice
	mySlice = append(mySlice, newslice...) // append newslice values to mySlice
	fmt.Println(mySlice)
}
