package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(sal Salutation, do Printer, isFormal bool) {
	message, altmessage := CreateMessage(sal.Name, sal.Greeting)
	if isFormal {
		do(message)
	}
	do(altmessage)
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func PrintCustomLine(s string, custom string) {
	fmt.Println(s + custom)
}

func CreateMessage(name, greeting string) (message string, altmessage string) {
	// fmt.Println(len(greeting))
	message = greeting + " " + name
	altmessage = "Hey alternate " + name
	return
}
