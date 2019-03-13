package main

import (
	"fmt"
	"log"
)

const fileName = "CoinKeeper_export.csv"
const expensesLimit = 1893

func main() {

	//fmt.Printf("unique tags: %v\n", allExpenses)
	//fmt.Printf("unique tags: %#v\n", selectUniqueTags(allExpenses))

	allExpenses, err := getExpensesFromFile(fileName, expensesLimit)
	if err != nil {
		log.Fatalf("unable to get expenses: %v", err)
	}

	sum, err :=  allExpenses.sumAllForTag(targetTag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("всего потрачено в %s: %v", targetTag, sum)
}