package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func greet(sal Salutation) {
	message, altmessage := CreateMessage(sal.name, sal.greeting)
	fmt.Println(altmessage)
	fmt.Println(message)
}

func CreateMessage(name, greeting string) (message string, altmessage string) {
	message = greeting + " " + name
	altmessage = "Hey alternate " + name
	return
}

func main() {
	var salute = Salutation{}
	salute.name = "chx"
	salute.greeting = "hello there"
	greet(salute)
}
