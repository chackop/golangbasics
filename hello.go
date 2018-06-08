package main

import "golangbasics/greeting"

func main() {
	// var salute = greeting.Salutation{}
	slice := []greeting.Salutation{
		{"chx", "Hey"},
		{"dimps", "mhhmmm"},
		{"chacko", "hellos"},
	}
	greeting.Greet(slice, greeting.CreatePrintFunction("custom CreateMessage"), true, 5)

	// salute.Name = "Bob"
	// salute.Greeting = "hello there"
	// greeting.Greet(salute, greeting.CreatePrintFunction("custom CreateMessage"), true, 5)
	// greeting.TypeSwitchTest(salute)
}
