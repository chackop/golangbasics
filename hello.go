package main

import (
	"fmt"
	"golangbasics/greeting"
)

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {
	// var salute = greeting.Salutation{}
	// var sl []int
	// sl = make([]int, 3)
	// sl[0] = 1
	// sl[1] = 10
	// sl[2] = 200

	salutes := greeting.Salutations{
		{"chx", "Hey"},
		{"dimps", "mhhmmm"},
		{"chacko", "hellos"},
	}

	// slice = append(slice[:1], slice[2:]...)

	// slice = slice[1:]

	// greeting.Greet(slice, greeting.CreatePrintFunction("custom CreateMessage"), true, 5)

	// salutes[0].Rename("John")

	// RenameToFrog(&salutes[0])

	fmt.Fprintf(&salutes[0], "the count is %d", 10)
	salutes.Greet(greeting.CreatePrintFunction("custom CreateMessage"), true, 5)

	// salute.Name = "Bob"
	// salute.Greeting = "hello there"
	// greeting.Greet(salute, greeting.CreatePrintFunction("custom CreateMessage"), true, 5)
	// greeting.TypeSwitchTest(salute)
}
