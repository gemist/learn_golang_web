package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["alan"])
}

func changeMe(z map[string]int) {
	z["alan"] = 30
}
