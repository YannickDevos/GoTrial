package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name, s.Greeting)

		//message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
		/*	i := 0
			for {
				if i >= times {
					break
				}
				if i%2 == 0 { //if number is disible by 2, it'll restart the loop without executing what's following
					i++
					continue
				}*/

		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
		//i++
	}
}

func GetPrefix(name string) (prefix string) {

	var prefixmap string string
	switch {
	case name == "Bob":
		prefix = "Mr "
		//fallthrough
	case name == "Joe", name == "Amy", len(name) == 10:
		prefix = "Dr "
	case name == "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "
	}
	return
}

/*
func TypeSwitchTest(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknown")
	}
}
*/

func CreateMessage(name, greeting string) (message string, alternate string) {
	fmt.Println(len(greeting)) //Print the number of parameters in the Greeting's slice
	message = greeting + " " + name
	alternate = "Hey!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func PrintCustom(s string, custom string) {
	fmt.Println(s, PrintCustom)
}
