package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "yannick devos"
	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)   //convert to TitleCase
	author = strings.ToUpper(author) //convert to Capitrals

	return module, author //return output to the function's caller + close the function
}
