package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

const (
	PI       = 3.14
	Language = "Go"
	A        = iota // iota create an incremental sequence in the constant
	B        = iota
	C        = iota
)

const (
	D = iota // if not repeated, the constant will assume that ioata is the same value for the rest of the constant
	E
	G
)

func main() {

	message := "Hello Go World!"
	a, b, c := 2, true, 9
	var greeting *string = &message
	*greeting = "hi"
	var s = Salutation{"Joe", "Hello!"}

	fmt.Println(message, a, b, c, *greeting, s.name)
	fmt.Println(s.greeting)
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A, B, C)
	fmt.Println(D, E, G)

}
