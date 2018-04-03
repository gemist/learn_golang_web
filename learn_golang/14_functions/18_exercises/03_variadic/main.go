package main

import "fmt"

func main() {
	fmt.Println(average(-1, -2, -3, -4, -5, -6))
}

func average(input ...float64) float64 {
	largest := input[0]
	for _, value := range input {
		if value > largest {
			largest = value
		}
	}
	return largest
}
