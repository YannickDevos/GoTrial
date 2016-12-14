package main

import (
	"fmt"
)

func main() {
	// declares a slice called mycourses
	//mycourses := make([]string, 5, 10)                  //create a slice of strings with initial length of 5 and a capacity of ten
	mycourses := []string{"Docker", "Puppet", "Python"} //create a slice of strings with initial length of 5 and a capacity of 3

	fmt.Printf("length is: %d. \nCapacity is: %d.",
		len(mycourses), cap(mycourses))
}
