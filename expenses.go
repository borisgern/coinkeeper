package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Payment struct {
	Data string
	Type string
	From string
	To string
	Tags []string
	Amount string
}

const targetTag = "вкусвилл"

type expenses []Payment

func getExpensesFromFile(fileName string, limit int) (expenses, error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	linesLimit := len(lines)

	if limit > linesLimit {
		return nil, fmt.Errorf("max limit %q", )
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

	return allExpenses, nil
}

func (exp expenses) selectUniqueTags() []string {
	uniqueTags := make(map[string]struct{},0)
	for _, e := range exp {
		//fmt.Printf("tags: %#v\n", s)
		appendUnique(uniqueTags, e.Tags)
	}
	return convertTags(uniqueTags)
}

func appendUnique(uniqueTags map[string]struct{}, tags []string) map[string]struct{} {
	for _, t := range tags {
		if _, ok := uniqueTags[t]; !ok {
			uniqueTags[t] = struct{}{}
		}
	}
	return uniqueTags
}

func convertTags(uniqueTags map[string]struct{}) []string {
	tags := make([]string,0,len(uniqueTags))
	for t := range uniqueTags {
		tags = append(tags, t)
	}
	return tags
}

func (exp expenses) sumAllForTag(tag string) (float64, error) {
	var sum float64
	for _, e := range exp {
		for _, t := range e.Tags {
			if t == tag {
				amount, err := strconv.ParseFloat(e.Amount,2)
				if err != nil {
					return 0, fmt.Errorf("unable to parse float: %v", err)
				}
				sum += amount
			}
		}
	}
	return sum, nil
}