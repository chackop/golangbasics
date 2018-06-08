package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func greet(sal Salutation) {
	message, altmessage := CreateMessage(sal.name, sal.greeting, "yo")
	fmt.Println(altmessage)
	fmt.Println(message)
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
	greet(salute)
}
