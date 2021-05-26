package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	n := flag.Bool("n", false, "omit trailing newline")
	sep := flag.String("s", " ", "separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}
