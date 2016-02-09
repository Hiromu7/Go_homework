// ren_1_3 project ren_1_3.go
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	var s, sep string
	start_roop := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Printf(s)
	time_roop := time.Since(start_roop).Seconds()
	fmt.Printf(" %.4fs elapsed\n", time_roop)

	start_join := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.4fs elapsed\n", time.Since(start_join).Seconds())
}
