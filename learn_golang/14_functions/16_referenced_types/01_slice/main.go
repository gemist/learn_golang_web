package main

import "fmt"

func main() {
	m := make([]string, 2, 8)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "alan"
	z[1] = "ford"
	fmt.Println(z)
}
