package main

import (
	"fmt"
)

type Yen float64
type Dollar float64

func YenToDollar(y Yen) Dollar {
	return Dollar(y / 120)
}

func DollarToYen(d Dollar) Yen {
	return Yen(d * 120)
}

func (y Yen) String() string { return fmt.Sprintf("%.2fYen", y) }

func (d Dollar) String() string { return fmt.Sprintf("%.2fDollar", d) }
