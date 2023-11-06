package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
    ZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

func FahrenheitToCelsius (f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

var cel = flag.Float64("Celsius", 0, "Celsius value to be converted to Fahrenheit")
var fahr = flag.Float64("Fahrenheit", 0, "Fahrenheit value to be converted to Celsius")
func main() {
	flag.Parse();
	var sampleC Celsius = 100

	fmt.Println(" %g ", CelsiusToFahrenheit(sampleC))

	if *cel != 0 {
		fmt.Println("\nInput Cel = %v, Fahr = %v", *cel, CelsiusToFahrenheit(Celsius(*cel)))
	}

	if *fahr != 0 {
		fmt.Println("\nInput Fahr = %v, Cel = %v", *fahr, FahrenheitToCelsius(Fahrenheit(*fahr)))
	}
}