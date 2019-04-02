package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"test/coinkeeper/proto"
	"test/coinkeeper/util"
)

func main() {
	conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}

	defer conn.Close()

	client := expensespb.NewExpensesServiceClient(conn)

	payments, err := getExpenses(client, "9/23/2018", "3/12/2019", 5)
	if err != nil {
		log.Fatalf("unable to get expenses: %v", err)
	}
	log.Printf("payments: %#v\n", payments[0])
}

func getExpenses(client expensespb.ExpensesServiceClient, fromDate, toDate string, limit int32) ([]*expensespb.Payment, error) {
	log.Printf("New GetExpenses request: from %q to %q with limit %v", fromDate, toDate, limit)
	ctx := context.Background()

	unixFromDate, err := util.ToUnixFormat(fromDate)
	if err != nil {
		return nil, fmt.Errorf("unable to convert date: %v", err)
	}

	unixToDate, err := util.ToUnixFormat(toDate)
	if err != nil {
		return nil, fmt.Errorf("unable to convert date: %v", err)
	}

	req := &expensespb.ExpensesRequest{
		FromDate:unixFromDate,
		ToDate:unixToDate,
		Limit:limit,
	}

	res, err := client.GetExpenses(ctx, req)
	return res.Payments, err
}