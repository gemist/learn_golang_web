package main

import "fmt"

func main() {
	n := 600851475143
	//n := 13195
	fmt.Println(largest_prime(n))
}

func largest_prime(n int) int {
	largest_prime := 1
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			largest_prime = i
			n = n / i
			i = i - 1
		}
	}

	largest_prime = n

	return largest_prime
}

/*The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
