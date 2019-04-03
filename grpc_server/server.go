package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"test/coinkeeper/proto"
	"test/coinkeeper/services"
	"test/coinkeeper/util"
)

const fileName = "./../CoinKeeper_export.csv"

type ExpensesManager struct {
	Logger *services.CtxLogger
}

func (em *ExpensesManager) GetExpenses(ctx context.Context, req *expensespb.ExpensesRequest) (*expensespb.Expenses, error) {
	log.Printf("GetExpenses function request: %v", req)
	limit := req.GetLimit()

	f, err := os.Open(fileName)
	if err != nil {
		return &expensespb.Expenses{}, fmt.Errorf("unable to open file: %v", err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	linesLimit := len(lines)

	if int(limit) > linesLimit {
		return &expensespb.Expenses{}, fmt.Errorf("max limit %v", linesLimit)
	}

	if limit == 0 {
		limit = int32(linesLimit)
	}

	fmt.Printf("number of lines %v, limit %v\n", linesLimit, limit)
	allExpenses := make([]*expensespb.Payment, limit)
	for i, line := range lines {
		em.Logger.Debugf("pidr date %v", line[0])
		date, err := util.ToUnixFormat(line[0])
		if err != nil {
			return &expensespb.Expenses{}, fmt.Errorf("unable to parse date: %v", err)
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