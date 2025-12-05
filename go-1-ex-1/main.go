package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	var firstName string = "Zainab"
	var lastName string = "Kalejaiye"

	dayOfBirth := 1
	monthOfBirth := 1
	yearOfBirth := 2009

	var numberOfSiblings = 2

	heightInMeters := 1.75

	var zodiacSign rune = '\u2651'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
