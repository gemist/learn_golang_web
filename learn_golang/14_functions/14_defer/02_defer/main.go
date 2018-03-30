package main

import "fmt"

func world() {
	fmt.Println("world")
}

func hello() {
	fmt.Println("Hello")
}

func main() {
	defer world()
	hello()
}
