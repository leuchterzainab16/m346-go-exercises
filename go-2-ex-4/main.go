package main

import (
	"fmt"
)

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class map[byte]Student
	classA := Class{
		1: Student{
			FirstName: "Joe",
			LastName:  "Doe",
		},
		2: Student{
			FirstName: "Jay",
			LastName:  "Day",
		},
	}
	classB := Class{
		3: Student{
			FirstName: "Sara",
			LastName:  "Haas",
		},
		4: Student{
			FirstName: "Jenny",
			LastName:  "Smith",
		},
	}

	classC := Class{
		3: Student{
			FirstName: "Hans",
			LastName:  "Gritch",
		},
		4: Student{
			FirstName: "Gretel",
			LastName:  "Gritch",
		},
	}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[string]Class{
		"M114": classA,
		"M346": classB,
		"M117": classC,
	}
	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
