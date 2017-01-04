/**
* Created by
* User:
* Date
* Time
* To change this template use File|Settings|File templates
**/

package main

import "./greeting"

func main() {

	s := greeting.Salutation{"ttttttnitp", "Hello!"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true)
	greeting.TypeSwitchTest("hello")

}
