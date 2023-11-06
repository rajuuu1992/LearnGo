// Package congerting Temperature units between Celsius, Fahrenheit, Kelvin)
package temperatureConverter

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	ZeroC     Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

// func (c Celsius) String() string {
//     return fmt.Sprintf("%g *C", c);
// }

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g *C", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g *K", k)
}
