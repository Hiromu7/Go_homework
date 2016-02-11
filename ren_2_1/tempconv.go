// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func CtoK(c Celsius) Kelvin         { return Kelvin(c - AbsoluteZeroC) }
func FtoC(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }
func FtoK(f Fahrenheit) Kelvin      { return CtoK(FtoC(f)) }
func KtoC(k Kelvin) Celsius         { return Celsius(k) + AbsoluteZeroC }
func KtoF(k Kelvin) Fahrenheit      { return CtoF(KtoC(k)) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

//!-
