package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//1
	//fmt.Println("Hello Go")

	//2
	/*var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)*/

	//3
	/*s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s)*/

	//4
	//fmt.Println(strings.Join(os.Args[1:], " "))

	//zad 1
	//fmt.Println(strings.Join(os.Args[:0], " "))

	//zad 2 i 3
	/*start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		fmt.Print(time.Since(start).Seconds())
		fmt.Print(" ")
		fmt.Print(i)
		fmt.Println(" " + os.Args[i])
	}*/

	//dup1
	/*counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}*/

	//dup2
	/*counts2 := make(map[string]int)
		files := os.Args[1:]

		if len(files) == 0 {
			countLines(os.Stdin, counts2)
		} else {
			for _, arg := range files {
				f, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
					continue
				}
				countLines(f, counts2)
				f.Close()
			}
		}
		for line, n := range counts2 {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}

	func countLines(f *os.File, counts map[string]int) {
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}
	}*/

	//dup3
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
