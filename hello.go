package main

import "golangbasics/greeting"

func main() {
	// var salute = greeting.Salutation{}
	// var sl []int
	// sl = make([]int, 3)
	// sl[0] = 1
	// sl[1] = 10
	// sl[2] = 200

	slice := []greeting.Salutation{
		{"chx", "Hey"},
		{"dimps", "mhhmmm"},
		{"chacko", "hellos"},
	}

	slice = append(slice[:1], slice[2:]...)

	// slice = slice[1:]

	greeting.Greet(slice, greeting.CreatePrintFunction("custom CreateMessage"), true, 5)

	// salute.Name = "Bob"
	// salute.Greeting = "hello there"
	// greeting.Greet(salute, greeting.CreatePrintFunction("custom CreateMessage"), true, 5)
	// greeting.TypeSwitchTest(salute)
}
