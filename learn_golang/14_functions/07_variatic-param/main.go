package main

import "fmt"

func main() {
	n := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
