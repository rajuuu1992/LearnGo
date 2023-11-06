package main

import (
	"temperatureConverter"
	"fmt"
	"flag"
)

var cel = flag.Float64("Celsius", 0, "Celsius value to be converted to Fahrenheit")
var fahr = flag.Float64("Fahrenheit", 0, "Fahrenheit value to be converted to Celsius")
var kelv = flag.Float64("Kelvin", 0, "Kelvin value to be converted to Celsius")

func main() {
	flag.Parse();
	var sampleC temperatureConverter.Celsius = 100

	fmt.Println(" %g ", temperatureConverter.CelsiusToFahrenheit(sampleC))

	if *cel != 0 {
		fmt.Println("\nInput Cel = %v, Fahr = %v", *cel, temperatureConverter.CelsiusToFahrenheit(temperatureConverter.Celsius(*cel)))
	}

	if *fahr != 0 {
		fmt.Println("\nInput Fahr = %v, Cel = %v", *fahr, temperatureConverter.FahrenheitToCelsius(temperatureConverter.Fahrenheit(*fahr)))
	}

	if *kelv != 0 {
		fmt.Println("\nInput Kelving = %v, Cel = %v", *kelv, temperatureConverter.KelvinToCelsius(temperatureConverter.Kelvin(*kelv)))
	}

	fmt.Println(" Constant %g", temperatureConverter.ZeroC)
}