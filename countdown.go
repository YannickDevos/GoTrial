package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- { //start at 10; loop operate while >=0; minus 1 on each iteration (decremen t the time by 1)
		if timer == 0 { //if timer hit 0, print boom
			fmt.Println("Boom!")
			break //break out of the current loop once Boom is displayed
		}
		fmt.Println(timer)          // Print the number of second until reach 0
		time.Sleep(1 * time.Second) //insert 1 second sleep on each loop iteration
	}
}
