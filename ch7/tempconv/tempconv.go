package tempconv

import (
	"flag"
	"fmt"
)

type Celcius float64
type Fahrenheit float64

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32.0) * 5.0 / 9.0) }

func (c Celcius) String() string { return fmt.Sprintf("%g°C", c) }

type celciusFlag struct{ Celcius }

func (f *celciusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celcius = Celcius(value)
		return nil
	case "F", "°F":
		f.Celcius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("Invalid temprature %q", s)
}

func CelciusFlag(name string, value Celcius, usage string) *Celcius {
	f := celciusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celcius
}
