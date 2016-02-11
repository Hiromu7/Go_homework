// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputs []string = os.Args[1:]

	if len(inputs) == 0 {
		var value string
		fmt.Scan(&value)
		inputs = strings.Split(value, " ")
	}

	for _, arg := range inputs {
		input, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		printconv(input)
	}
}

func printconv(input float64) {
	m := Meter(input)
	s := Shaku(input)
	fmt.Printf("%s = %s, %s = %s\n", m, MeterToShaku(m), s, ShakuToMeter(s))

	y := Yen(input)
	d := Dollar(input)
	fmt.Printf("%s = %s, %s = %s\n", d, DollarToYen(d), y, YenToDollar(y))
}

//!-
