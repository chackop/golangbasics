package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func Greet(sal Salutation, do Printer) {
	message, altmessage := CreateMessage(sal.name, sal.greeting, "yo")
	do(message)
	do(altmessage)
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

func PrintCustomLine(s string, custom string) {
	fmt.Println(s + custom)
}

func CreateMessage(name string, greeting ...string) (message string, altmessage string) {
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	altmessage = "Hey alternate " + name
	return
}

func main() {
	var salute = Salutation{}
	salute.name = "chx"
	salute.greeting = "hello there"
	Greet(salute, createPrintFunction("custom CreateMessage"))
}
