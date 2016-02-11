// ren_1_4 project ren_1_4.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %c\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)

			for _, arg := range files {
				f, err := os.Open(arg)
				if err == nil {
					word := bufio.NewScanner(f)
					for word.Scan() {
						if word.Text() == line {
							fmt.Print(arg)
							fmt.Printf("\t")
						}
					}
					f.Close()
				}
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
