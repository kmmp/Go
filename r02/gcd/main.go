package main

import (
	"fmt"
)

func main() {
	x := 25
	y := 50

	for y != 0 {
		x, y = y, x%y
		fmt.Printf("%d\t%d\n", x, y)
	}
}
