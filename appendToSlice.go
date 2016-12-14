package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is: %d Capacity is: %d",
		len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ { //append 1 int at a time to the array until it reach 17
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", // check the array capacity evolution after each increment (array double its capacity each time it reach its max value)
			cap(mySlice))
	}
}
