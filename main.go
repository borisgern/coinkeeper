package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
	"test/coinkeeper/proto"
)

const fileName = "CoinKeeper_export.csv"
const expensesLimit = 1893
const serverHost = "0.0.0.0"
const serverPort = "50051"

type ExpensesManager struct {

}

func main() {

	//fmt.Printf("unique tags: %v\n", allExpenses)
	//fmt.Printf("unique tags: %#v\n", selectUniqueTags(allExpenses))

	lis, err := net.Listen("tcp", serverHost + ":" + serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	expensespb.RegisterExpensesServiceServer(server, &ExpensesManager{})

	allExpenses, err := getExpensesFromFile(fileName, expensesLimit)
	if err != nil {
		log.Fatalf("unable to get expenses: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	sum, err :=  allExpenses.sumAllForTag(targetTag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("всего потрачено в %s: %v", targetTag, sum)
}

func (em *ExpensesManager) GetExpenses(ctx context.Context, req *expensespb.ExpensesRequest) (*expensespb.Expenses, error) {
	log.Printf("GetExpenses function request: %v", req)
	limit := req.GetLimit()

	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	linesLimit := len(lines)

	if int(limit) > linesLimit {
		return nil, fmt.Errorf("max limit %q", linesLimit)
	}

	fmt.Printf("number of lines %v\n", linesLimit)

	allExpenses := make(expenses, limit)
	for i, line := range lines {
		tags := strings.Split(line[4], ", ")
		data := Payment{
			Data:line[0],
			Type:line[1],
			From:line[2],
			To: line[3],
			Tags:tags,
			Amount:line[5],
		}
		allExpenses[i] = data
		if i == limit-1 {
			break
		}
	}
}