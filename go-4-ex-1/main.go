package main

import "fmt"

func computeGrade(gotPoints float64, maxPoints float64) float64 {

	return (gotPoints/maxPoints)*5 + 1
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(45, 60)) // 4.75
	fmt.Println(computeGrade(30, 60)) // 3.5
	fmt.Println(computeGrade(15, 60)) // 2.25
}
