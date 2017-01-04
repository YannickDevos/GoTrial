package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	do(message)
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting)) //Print the number of parameters in the greeting's slice
	message = greeting[1] + " " + name
	alternate = "Hey!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func createPrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func PrintCustom(s string, custom string) {
	fmt.Println(s, PrintCustom)
}

func main() {

	s := Salutation{"Bob", "Hello!"}
	Greet(s, createPrintFunction("!!!"))

}
