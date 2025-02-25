/*
Assignment 2: Temperature Converter
Use these: Methods, Receivers, Pointers
Problem Statement: Create a program that converts temperatures between Celsius and Fahrenheit.
Requirements:
1. Define a struct Temperature with a Value (float64).
2. Implement methods:
  - ToFahrenheit() → Converts from Celsius to Fahrenheit.
  - ToCelsius() → Converts from Fahrenheit to Celsius.

3. Use pointers to modify the temperature directly.
Formula:
* °F = (°C × 9/5) + 32
* °C = (°F - 32) × 5/9
*/
package main

import "fmt"

// Temperature struct
type Temperature struct {
	Temp float64
}

// Method to convert Celsius to Fahrenheit
func (t *Temperature) ToFahrenheit() {
	t.Temp = (t.Temp * 9 / 5) + 32
}

// Method to convert Fahrenheit to Celsius, *Temp is using pointers
func (t *Temperature) ToCelsius() {
	t.Temp = (t.Temp - 32) * 5 / 9
}

func main() {
	//creating an instance of Temperature with temp 35
	//  Converting 35°C to Fahrenheit
	celsiusTemp := Temperature{Temp: 35}
	//Calls ToFahrenheit(), which converts 35°C to 95°F using the formula.
	celsiusTemp.ToFahrenheit()
	fmt.Printf("Converted: %.2f°F\n", celsiusTemp.Temp)

	//Creates an instance of Temperature with Temp: 95 (°F).
	//  Converting 95°F to Celsius
	fahrenheitTemp := Temperature{Temp: 95}
	//Calls ToCelsius(), which converts 95°F back to 35°C.
	fahrenheitTemp.ToCelsius()
	fmt.Printf("Converted: %.2f°C\n", fahrenheitTemp.Temp)
}
