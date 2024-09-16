package main

import "fmt"

func do(a [3]int, b []int) []int {
	// a = b    // SYNTAX ERROR
	a[0] = 4 // w unchanged
	b[0] = 3 // x changed because it is an alias of w, points to the same place in memory

	c := make([]int, 5) // []int{0, 0, 0, 0, 0}
	c[4] = 42
	copy(c, b) // copies only 3 elements

	return c
}

func main() {
	// Array
	var w = [...]int{1, 2, 3}

	// Slice
	var x = []int{0, 0, 0}

	y := do(w, x)

	fmt.Println(w, x, y) // [1 2 3] [3 0 0] [3 0 0 0 42]
}
