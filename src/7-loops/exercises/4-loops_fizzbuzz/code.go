package main

import (
	"fmt"
)

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		fizzOrBuzz := false
		if i%3 == 0 {
			fmt.Print("fizz")
			fizzOrBuzz = true
		}
		if i%5 == 0 {
			fmt.Print("buzz")
			fizzOrBuzz = true
		}
		if fizzOrBuzz {
			fmt.Print("\n")
		} else {
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
