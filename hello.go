package main

import (
	"fmt"
	"golangbasics/greeting"
)

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {
	// var salute = greeting.Salutation{}
	// var sl []int
	// sl = make([]int, 3)
	// sl[0] = 1
	// sl[1] = 10
	// sl[2] = 200

	salutes := greeting.Salutations{
		{"chx", "Hey"},
		{"dimps", "mhhmmm"},
		{"chacko", "hellos"},
	}

	// slice = append(slice[:1], slice[2:]...)

	// slice = slice[1:]

	// greeting.Greet(slice, greeting.CreatePrintFunction("custom CreateMessage"), true, 5)

	// salutes[0].Rename("John")

	// RenameToFrog(&salutes[0])

	fmt.Fprintf(&salutes[0], "the count is %d", 10)
	ch := make(chan greeting.Salutation)
	ch2 := make(chan greeting.Salutation)
	go salutes.ChannelGreeter(ch)
	go salutes.ChannelGreeter(ch2)

	for {
		select {
		case s, ok := <-ch:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}

		case s, ok := <-ch2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting")
		}

	}
	// for s := range ch {
	// 	fmt.Println(s.Name)
	// }
}
