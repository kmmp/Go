package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Meter float64
type Foot float64
type Pound float64
type Kg float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g Meter", m) }
func (f Foot) String() string       { return fmt.Sprintf("%g Foot", f) }
func (p Pound) String() string      { return fmt.Sprintf("%g Pound", p) }
func (k Kg) String() string         { return fmt.Sprintf("%g Kg", k) }
