package main

import (
	"fmt"
)

func main() {

	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //create a slice of strings with initial length of 10 and a capacity of 10

	fmt.Println(myslice[4]) //Print slice value at index # 4

	myslice[1] = 0
	fmt.Println(myslice)

	sliceofslice := myslice[2:5] // create a slice from orginal slice from index value 2 to 5 [WARNING: doesn't display end value of 5 though]
	fmt.Println(sliceofslice)
}
