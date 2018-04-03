package main

import "fmt"

func main() {
	i, b := half(5)
	fmt.Printf("%g, %t\n", i, b)
}

func half(i int) (float64, bool) {
	return float64(i) / 2, i%2 == 0
}
