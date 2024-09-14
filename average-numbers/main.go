package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var totalNumbers int

	for {
		var value float64

		_, err := fmt.Fscanln(os.Stdin, &value)

		if err != nil {
			// you can type control + D to end the program
			break
		}

		sum += value
		totalNumbers++
	}

	if totalNumbers == 0 {
		fmt.Fprintln(os.Stderr, "no values provided")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(totalNumbers))
}
