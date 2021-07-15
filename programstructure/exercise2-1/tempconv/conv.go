package tempconv

// Converts celsius to fahrenheight
func CToF(c Celcius) Fahrenheight { return Fahrenheight(c*9/5 + 32) }

// Converts fahrenhiehg to celcius
func FToC(f Fahrenheight) Celcius { return Celcius(f - 32*5/9) }

// Converts celsius to kelvin
func CToK(c Celcius) Kelvin { return Kelvin(Kelvin(c) + ZeroK) }

// Converts kelvin to celsius
func KToC(k Kelvin) Celcius { return Celcius(k + ZeroK) }
