package main

import "fmt"

func main() {

	mySlice := []string{"a", "b", "c", "d", "f", "g"}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[4])
	fmt.Println("myString"[2])

}
