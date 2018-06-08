package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(sal []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range sal {
		message, altmessage := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(altmessage)
		}
	}

}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "chx":
		prefix = "Mr "
		fallthrough
	case name == "Chacko", name == "chacs":
		prefix = "Dr "
	case name == "Dimps", len(name) == 10:
		prefix = "Mrs "
	default:
		prefix = "Hey all"
	}
	return
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("saluatation")
	default:
		fmt.Println("unknown")
	}
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
