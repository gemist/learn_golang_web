package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, I am here!")
	}()
}
