package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) (float64, float64) {
	discriminant := math.Pow(b, 2) - 4*a*c
	sqrtDiscriminant := math.Sqrt(discriminant)
	if discriminant < 0 {
		fmt.Println("Die Diskriminante ist negativ, es gibt keine reellen Wurzeln.")
	}
	root1 := (-b + sqrtDiscriminant) / (2 * a)
	root2 := (-b - sqrtDiscriminant) / (2 * a)
	return root1, root2
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // -0.3, -1
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // -1, -1
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // -2, -3
}
