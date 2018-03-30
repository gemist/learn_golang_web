package main

import "fmt"

func main() {
	age := 33
	fmt.Println(&age)
	fmt.Println(age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)

}

func changeMe(num *int) {
	fmt.Println(num)
	fmt.Println(*num)
	*num = 10
	fmt.Println(num)
	fmt.Println(*num)
}
