package main

import "fmt"

func main() {
	testmap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5} //create map (remember = UNORDERED list)
	for key, value := range testmap {                                 //loop through key/value of the map
		fmt.Printf("\nKey is: %v Value is: %v", key, value) //print key/value on different line (note that due to GO "random starting offset , the starting value will be different on every run)
	}

	testmap["A"] = 100  //Update A value to 100
	testmap["F"] = 1985 //add a new value to the map as F value doesn't exist in it yet
	fmt.Println(testmap)

	delete(testmap, "F") //delete BOTH key/value for F value
	fmt.Println(testmap)
}
