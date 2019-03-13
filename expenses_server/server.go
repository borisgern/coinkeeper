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
	"test/coinkeeper/util"
)

const fileName = "CoinKeeper_export.csv"
const serverHost = "0.0.0.0"
const serverPort = "50052"

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

	//allExpenses, err := getExpensesFromFile(fileName, expensesLimit)
	//if err != nil {
	//	log.Fatalf("unable to get expenses: %v", err)
	//}
	log.Printf("server listner started")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//sum, err :=  allExpenses.sumAllForTag(targetTag)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("всего потрачено в %s: %v", targetTag, sum)
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
	allExpenses := make([]*expensespb.Payment, limit)
	for i, line := range lines {
		date, err := util.ToUnixFormat(line[0])
		if err != nil {
			return nil, fmt.Errorf("unable to parse date: %v", err)
		}
		tags := strings.Split(line[4], ", ")
		data := expensespb.Payment{
			Date: date,
			Type:line[1],
			From:line[2],
			To: line[3],
			Tags:tags,
			Amount:line[5],
		}
		allExpenses[i] = &data
		if i == int(limit)-1 {
			break
		}
	}
	return &expensespb.Expenses{Payments:allExpenses}, nil
}