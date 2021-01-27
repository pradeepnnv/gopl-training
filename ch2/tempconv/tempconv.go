package tempconv

import "fmt"

//Celsius is a new type (whose underlying type is float64) which represents temperature in Celsius
type Celsius float64

//Fahrenheit is a new type (whose underlying type is float64) which represents temperature in Fahrenheit
type Fahrenheit float64

// AbsoluteZeroC, FreezingC, BoilingC represents various constants being used
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
