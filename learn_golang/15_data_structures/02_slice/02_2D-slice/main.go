package main

import "fmt"

func main() {

	records := make([][]string, 0)

	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "77.0"

	records = append(records, student1)

	fmt.Println(records)

	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.0"
	student2[3] = "96.0"

	records = append(records, student2)

	fmt.Println(records)

	// loop example

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)

		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println("transactions: ", transactions)
}
