package main

import "fmt"

func main() {
	half := func(i int) (float64, bool) {
		return float64(i) / 2, i%2 == 0
	}

	fmt.Println(half(10))
}
