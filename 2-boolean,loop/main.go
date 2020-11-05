package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Boolean
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave: ", exit)

	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	var age = 41
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	fmt.Println("There is a cavern entrance here and a path to the east.")
	var instruction = "go inside"

	// Switch
	switch instruction {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	default:
		fmt.Println("Didn't quite get that.")
	}

	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}

	// Loop
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")

}
