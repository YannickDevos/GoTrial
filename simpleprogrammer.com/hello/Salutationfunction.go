package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {

	s := Salutation{"Bob", "Hello!"}
	Greet(s)

}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting)) //Print the number of parameters in the greeting's slice
	message = greeting[1] + " " + name
	alternate = "Hey!" + name
	return
}
