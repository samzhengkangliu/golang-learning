package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.\n")

	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be", 41*365.2425/687, "years old")

	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)

	// similar to interpolations
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)

	// -15 fill 15 spaces to the right
	// 4 fill 4 spaces to the left
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println(distance/lightSpeed, "seconds")

	var weight = 149.0
	// equivalent to weight = weight * 0.3783
	weight *= 0.3783
	var age = 25
	// equivalent to age = age + 1
	age += 1
	age++

	// [0, n)
	var num = rand.Intn(10) + 1
	fmt.Println(num)
}
