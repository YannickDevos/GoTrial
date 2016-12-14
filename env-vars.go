package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() { // will loop through os.Environ list and print one parameter per line
		fmt.Println(env)
	}
}
