package main

import "fmt"

func main() {
	message := "hello world"
	var greeting *string = &message
	*greeting = "hi message"
	fmt.Println(message, *greeting)
}
