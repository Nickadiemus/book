// Package tempconv performs Celsius and Fahrenheight temperature computations.
package tempconv

import (
	"flag"
	"fmt"
)

// delcaring custom types for utility
type Fahrenheit float64
type Celsius float64
type Kelvin float64

// type Delcarations
const (
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
	ZeroK     Kelvin  = -273.15
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

//7.4 addition *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	// Sscanf parses a floating point value (&value and a string &unit)
	// from input
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)

}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
