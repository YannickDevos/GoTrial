package main

import (
	"fmt"
	"os"
	"reflect"
)

// Package level variable declaration (= global)
var (
	name   = "Nigel"            //Name of the subscriber (Type String)
	course = "Docker Deep Dive" //Couse being viewed (Type String)
	module = 3.2                //Current position in course (Type float)
)

func main() {
	// Function level variable declaration + initialization (= work ONLY inside of the function)
	// Variables declared at fucntion level MUST be used (error otherwise)
	//Function level variable overwrite Package level variable
	name := os.Getenv("USERNAME") //retrieve name of the windows session user (Type String)
	//course := "Docker Deep Dive 2" //Couse being viewed (Type String)
	module := 3.4  //Current position in course (Type float)
	ptr := &module //declare pointer value of module directly

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memory address of *module* variable is", &module) //Print the memory hexadecimal address value of the variable instead of its type = aka pointer value
	fmt.Println("Memory address of *module* variable is", ptr,     //Do the same as previous using ptr value
		"and the value of *module* is", *ptr) //Use * to de-reference the pointer => display the variable's value
}
