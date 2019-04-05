package main

import (
	"context"
	"encoding/csv"
	"fmt"
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
	CSVLines [][]string
}

type CategoryData struct {
	Amount float32
	Frequency int32
}

func (em *ExpensesManager) GetExpenses(ctx context.Context, req *expensespb.ExpensesRequest) (*expensespb.Expenses, error) {
	em.Logger.Printf("GetExpenses function request: %v", req)
	limit := req.GetLimit()
	tag := req.GetTag()

	if len(em.CSVLines) == 0 {
		lines, err := readCSVFile()
		if err != nil {
			return &expensespb.Expenses{}, err
		}
		em.CSVLines = lines
	}

	linesLimit := len(em.CSVLines)

	if int(limit) > linesLimit {
		return &expensespb.Expenses{}, fmt.Errorf("max limit %v", linesLimit)
	}

	if limit == 0 {
		limit = int32(linesLimit)
	}

	em.Logger.Printf("number of lines %v, limit %v\n", linesLimit, limit)
	allExpenses := make([]*expensespb.Payment, 0, limit)
	for _, line := range em.CSVLines {
		date, err := util.ToUnixFormat(line[0])
		if err != nil {
			return &expensespb.Expenses{}, fmt.Errorf("unable to parse date: %v", err)
		}
		amount, err := strconv.ParseFloat(line[5], 2)
		if date >= req.GetFromDate() && date <= req.GetToDate() {
			tags := strings.Split(line[4], ", ")
			if len(tags) == 1 && tags[0] == "" {
				tags[0] = line[3]
			} else {
				tags = append(tags, line[3])
			}
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

func (em *ExpensesManager) GetCategories(ctx context.Context, req *expensespb.CategoriesRequest) (*expensespb.Categories, error) {
	em.Logger.Printf("GetCategories function request: %v", req)
	limit := req.GetLimit()

	if len(em.CSVLines) == 0 {
		lines, err := readCSVFile()
		if err != nil {
			return &expensespb.Categories{}, err
		}
		em.CSVLines = lines
	}

	linesLimit := len(em.CSVLines)

	if int(limit) > linesLimit {
		return &expensespb.Categories{}, fmt.Errorf("max limit %v", linesLimit)
	}

	if limit == 0 {
		limit = int32(linesLimit)
	}

	em.Logger.Printf("number of lines %v, limit %v\n", linesLimit, limit)
	allExpenses := make([]*expensespb.Payment, 0, limit)
	for _, line := range em.CSVLines {
		date, err := util.ToUnixFormat(line[0])
		if err != nil {
			return &expensespb.Categories{}, fmt.Errorf("unable to parse date: %v", err)
		}
		amount, err := strconv.ParseFloat(line[5], 2)
		if date >= req.GetFromDate() && date <= req.GetToDate() {
			tags := strings.Split(line[4], ", ")
			if len(tags) == 1 && tags[0] == "" {
				tags[0] = line[3]
			} else {
				tags = append(tags, line[3])
			}
			data := expensespb.Payment{
				Date: date,
				Type:line[1],
				From:line[2],
				To: line[3],
				Tags:tags,
				Amount:float32(amount),
			}
			if data.Type == "Expense" {
				allExpenses = append(allExpenses, &data)
			}
		}

		if len(allExpenses) == int(limit) {
			break
		}
	}

	return &expensespb.Categories{Categories:getCategoriesWithAmount(allExpenses, int(limit))}, nil
}

func checkTag(tags []string, tag string) (ok bool) {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return
}

func readCSVFile() ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer f.Close()

	return csv.NewReader(f).ReadAll()
}

func  getCategoriesWithAmount(exp []*expensespb.Payment, limit int) []*expensespb.Category {
	uniqueTags := make(map[string]CategoryData,0)
	for _, e := range exp {
		for _, t := range e.Tags {
			tempData := CategoryData{
				Amount: uniqueTags[t].Amount + e.Amount,
				Frequency: uniqueTags[t].Frequency + 1,
			}
			uniqueTags[t]= tempData
		}
	}
	categoriesLen := limit
	if len(uniqueTags) < limit {
		categoriesLen = len(uniqueTags)
	}
	categories := make([]*expensespb.Category, categoriesLen)
	var i int
	for k, v := range uniqueTags {
		category := &expensespb.Category{
			Amount: v.Amount,
			Frequency: v.Frequency,
			Name: k,
		}
		categories[i] = category
		i++
		if i == limit {
			break
		}
	}
	return categories
}