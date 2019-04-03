package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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
	tag := req.GetTag()

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
	allExpenses := make([]*expensespb.Payment, 0, limit)
	for _, line := range lines {
		date, err := util.ToUnixFormat(line[0])
		if err != nil {
			return &expensespb.Expenses{}, fmt.Errorf("unable to parse date: %v", err)
		}
		amount, err := strconv.ParseFloat(line[5], 2)
		if date >= req.GetFromDate() && date <= req.GetToDate() {
			tags := strings.Split(line[4], ", ")
			data := expensespb.Payment{
				Date: date,
				Type:line[1],
				From:line[2],
				To: line[3],
				Tags:tags,
				Amount:float32(amount),
			}
			if (checkTag(data.Tags,tag) || tag == "") && data.Type == "Expense" {
				allExpenses = append(allExpenses, &data)
			}
		}

		if len(allExpenses) == int(limit) {
			break
		}
	}
	return &expensespb.Expenses{Payments:allExpenses}, nil
}

func checkTag(tags []string, tag string) (ok bool) {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return
}