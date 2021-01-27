package tempconv

// CToF converts temperature in Celsius to Fahrenheit. It returns Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts temperature in Fahrenheit to Celsius. It returns Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
