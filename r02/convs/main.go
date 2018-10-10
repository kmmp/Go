package main

import (
	"fmt"
	"os"
	"strconv"

	"../tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			m := tempconv.Meter(t)
			foot := tempconv.Foot(t)
			p := tempconv.Pound(t)
			k := tempconv.Kg(t)

			fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c), m, tempconv.MToF(m), foot, tempconv.FToM(foot), p, tempconv.PToK(p), k, tempconv.KToP(k))
		}
	} else {
		fmt.Println("Brak parametrów!")
	}
}
