package main

import "fmt"

func main() {
	fmt.Print(greet("Alan", "Bizjak"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
