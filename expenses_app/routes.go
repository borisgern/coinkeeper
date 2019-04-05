package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"test/coinkeeper/proto"
	. "test/coinkeeper/services"
	"test/coinkeeper/util"
	"sort"
)

func(app *expensesApp) newRouter(config *HTTPServerConfig, log *CtxLogger) *gin.Engine {
	if config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ctxLog := log.NewPrefix("router")
	router := gin.New()

	router.RedirectTrailingSlash = false
	router.Use(RequestIDMiddleware(), LogMiddleware(ctxLog))
	router.Delims("{{", "}}")
	router.LoadHTMLGlob("./../templates/*.html")

	router.GET("/", app.RenderMain)

	expensesGroup := router.Group("/expenses")
	expensesGroup.GET("/", app.GetExpenses)

	categoriesGroup := router.Group("/categories")
	categoriesGroup.GET("/", app.GetCategories)

	return router
}

func(app *expensesApp) RenderMain(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "all_expenses.html",nil)
}

func(app *expensesApp) GetCategories(ctx *gin.Context) {
	limit := ctx.Query("limit")
	fromDate := ctx.Query("from")
	toDate := ctx.Query("to")
	sorting := ctx.Query("sort")

	conn, err := grpc.Dial(app.Config.GRPCServer.ServerHost+":"+app.Config.GRPCServer.ServerPort, grpc.WithInsecure())
	if err != nil {
		ErrorBadRequest(ctx, fmt.Errorf("couldn't connect: %v", err), "")
		return
	}

	defer conn.Close()

	client := expensespb.NewExpensesServiceClient(conn)

	var limitNum int
	if limit != "" {
		limitNum, err = strconv.Atoi(limit)
		if err != nil {
			ErrorBadRequest(ctx, fmt.Errorf("wrong limit format: %v", err), "")
			return
		}
	}

	categories, err := app.getCategories(client, fromDate, toDate, limitNum)
	if err != nil {
		err = fmt.Errorf("unable to get categories: %v", err)
	}

	if sorting == "true" {
		sort.Slice(categories, func(i, j int) bool { return categories[i].Amount > categories[j].Amount})
	}
	HandleResultOrError(ctx, gin.H{"categories": categories, "errorMsg": err}, err, "")
}

func(app *expensesApp) GetExpenses(ctx *gin.Context) {
	limit := ctx.Query("limit")
	fromDate := ctx.Query("from")
	toDate := ctx.Query("to")
	tag := ctx.Query("tag")
	sorting := ctx.Query("sort")

	conn, err := grpc.Dial(app.Config.GRPCServer.ServerHost+":"+app.Config.GRPCServer.ServerPort, grpc.WithInsecure())
	if err != nil {
		ErrorBadRequest(ctx, fmt.Errorf("couldn't connect: %v", err), "")
		return
	}

	defer conn.Close()

	client := expensespb.NewExpensesServiceClient(conn)

	var limitNum int
	if limit != "" {
		limitNum, err = strconv.Atoi(limit)
		if err != nil {
			ErrorBadRequest(ctx, fmt.Errorf("wrong limit format: %v", err), "")
			return
		}
	}


	payments, err := app.getExpenses(client, fromDate, toDate, tag, limitNum)
	if err != nil {
		err = fmt.Errorf("unable to get expenses: %v", err)
	}

	var tagSum float32
		tagSum = app.sumAllForTag(payments, tag)
		if err != nil {
			app.Logger.Errorf("unable to sum amount for tag %v: %v", tag, err)
	}

	if sorting == "true" {
		sort.Slice(payments, func(i, j int) bool { return payments[i].Amount > payments[j].Amount})
	}

	HandleResultOrError(ctx, gin.H{"expenses": payments, "tagSum" : tagSum, "errorMsg": err}, err, "")
}

func(app *expensesApp) getExpenses(client expensespb.ExpensesServiceClient, fromDate, toDate, tag string, limit int) ([]*expensespb.Payment, error) {
	app.Logger.Printf("New GetExpenses request: from %q to %q with limit %v", fromDate, toDate, limit)
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
		Limit:int32(limit),
		Tag:tag,
	}

	res, err := client.GetExpenses(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Payments, nil
}

func(app *expensesApp) getCategories(client expensespb.ExpensesServiceClient, fromDate, toDate string, limit int) ([]*expensespb.Category, error) {
	app.Logger.Printf("New GetCategories request: from %q to %q with limit %v", fromDate, toDate, limit)
	ctx := context.Background()

	unixFromDate, err := util.ToUnixFormat(fromDate)
	if err != nil {
		return nil, fmt.Errorf("unable to convert date: %v", err)
	}

	unixToDate, err := util.ToUnixFormat(toDate)
	if err != nil {
		return nil, fmt.Errorf("unable to convert date: %v", err)
	}

	req := &expensespb.CategoriesRequest{
		FromDate:unixFromDate,
		ToDate:unixToDate,
		Limit:int32(limit),
	}

	res, err := client.GetCategories(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Categories, nil
}

func(app *expensesApp) sumAllForTag(expenses []*expensespb.Payment, tag string) float32 {
	var sum float32
	for _, e := range expenses {
		if tag != "" {
			for _, t := range e.Tags {
				if t == tag {
					sum += e.Amount
					break
				}
			}
		} else {
			sum += e.Amount
		}
	}
	return sum
}