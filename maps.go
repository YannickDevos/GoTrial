package main

import (
	"fmt"
)

func main() {
	leagueTitles := make(map[string]int) // Declare a map of strings (key values) and int (data values)
	leagueTitles["Sunderlands"] = 6      // initialize the map with data
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{ //alternative way to declare maps
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\nLeague titles: %v \nRecent head to heads: %v \n", // print maps
		leagueTitles, recentHead2Head)
}
