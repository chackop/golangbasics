package main

import "golangbasics/greeting"

func main() {
	var salute = greeting.Salutation{}
	salute.Name = "obmaryjoeq"
	salute.Greeting = "hello there"
	greeting.Greet(salute, greeting.CreatePrintFunction("custom CreateMessage"), true)
}
