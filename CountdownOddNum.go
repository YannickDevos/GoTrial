package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- { //start at 10; loop operate while >=0; minus 1 on each iteration (decremen t the time by1)
		if timer%2 == 0 { //Print odd numbers only
			continue
		}
		fmt.Println(timer)          // Print the number of second until reach 0
		time.Sleep(1 * time.Second) //insert 1 second sleep on each loop iteration
	}
}
