package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	Born             BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			FirstName: "Zainab",
			LastName:  "Kalejaiye",
		},
		Born: BirthDate{
			DayOfBirth:   1,
			MonthOfBirth: 1,
			YearOfBirth:  2009,
		},
		NumberOfSiblings: 2,        // TODO: adjust
		ZodiacSign:       '\u2651', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	// me.NumberOfSiblings = 3 (alternative)
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
