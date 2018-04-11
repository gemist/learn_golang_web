package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)
	fmt.Println("--------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------")

	for i := 0; i < 13; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("slice value: ", mySlice[i], " len: ", len(mySlice), " cap: ", cap(mySlice))

	}

	//append slice to a slice
	fmt.Println("initial slice: ", mySlice)
	mySlice = append(mySlice, mySlice...)
	fmt.Println("final   slice: ", mySlice)

	// delete from slice
	mySlice = append(mySlice[:2], mySlice[20:]...)
	fmt.Println("delete from slice :", mySlice)

}
