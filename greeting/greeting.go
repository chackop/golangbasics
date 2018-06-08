package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(sal Salutation, do Printer, isFormal bool) {
	message, altmessage := CreateMessage(sal.Name, sal.Greeting)
	if prefix := GetPrefix(sal.Name); isFormal {
		do(prefix + message)
	} else {
		do(altmessage)
	}
}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr "
		fallthrough
	case name == "Joe", name == "chx":
		prefix = "Dr "
	case name == "Mary", len(name) == 10:
		prefix = "Mrs "
	default:
		prefix = "Hey all"
	}
	return
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
