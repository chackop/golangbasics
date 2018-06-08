package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

const (
	pi       = 3.14
	language = "GO"
)

const (
	A = iota
	B
	C
)

func main() {
	var salute = Salutation{}
	salute.name = "chx"
	salute.greeting = "hello chx"
	fmt.Println(salute.name, salute.greeting)
	fmt.Println(pi, language)
	fmt.Println(A, B, C)
}
