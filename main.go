package main

import (
	"fmt"

	"example.com/m/greetings"
	"example.com/m/hellos"
)

func main() {
	fmt.Println("hello forme Jan Erik")
	fmt.Println("Helloooo from Tuva")
	fmt.Println("Aaaaaa from Karin")
	fmt.Println(greetings.HellosFromAll())
	fmt.Println(greetings.ByeToAll())
	fmt.Println(hellos.Tester())
}
