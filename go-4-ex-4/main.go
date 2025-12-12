package main

import (
	"fmt"
)

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	fmt.Println(convertCelsiusToFahrenheit(0))   // 32
	fmt.Println(convertCelsiusToFahrenheit(100)) // 212
	fmt.Println(convertCelsiusToFahrenheit(-40)) // -40

	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println(convertFahrenheitToCelsius(32))  // 0
	fmt.Println(convertFahrenheitToCelsius(212)) // 100
	fmt.Println(convertFahrenheitToCelsius(-40)) // -40
}
