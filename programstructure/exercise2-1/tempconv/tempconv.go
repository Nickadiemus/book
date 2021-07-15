// Package tempconv performs Celsius and Fahrenheight temperature computations.
package tempconv

import "fmt"

// delcaring custom types for utility
type Fahrenheight float64
type Celcius float64
type Kelvin float64

// type Delcarations
const (
	FreezingC Celcius = 0
	BoilingC  Celcius = 100
	ZeroK     Kelvin  = -273.15
)

func (c Celcius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheight) String() string { return fmt.Sprintf("%g°F", f) }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
