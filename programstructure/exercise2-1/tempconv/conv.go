package tempconv

// Converts celsius to fahrenheight
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// Converts fahrenhiehg to celcius
func FToC(f Fahrenheit) Celsius { return Celsius(f - 32*5/9) }

// Converts celsius to kelvin
func CToK(c Celsius) Kelvin { return Kelvin(Kelvin(c) + ZeroK) }

// Converts kelvin to celsius
func KToC(k Kelvin) Celsius { return Celsius(k + ZeroK) }
