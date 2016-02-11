package main

import (
	"fmt"
)

type Meter float64
type Shaku float64

func ShakuToMeter(s Shaku) Meter {
	return Meter(s * 0.3030)
}

func MeterToShaku(m Meter) Shaku {
	return Shaku(m / 0.3030)
}

func (m Meter) String() string { return fmt.Sprintf("%.2fm", m) }

func (s Shaku) String() string { return fmt.Sprintf("%.2fshaku", s) }
