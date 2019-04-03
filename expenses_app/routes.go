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
	router.LoadHTMLGlob("./../templates/*.tmpl.html")

	router.GET("/", app.RenderMain)

	expensesGroup := router.Group("/expenses")
	expensesGroup.GET("/", app.GetExpenses)

	return router
}

func(app *expensesApp) RenderMain(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl.html",nil)
}

func(app *expensesApp) GetExpenses(ctx *gin.Context) {
	limit := ctx.Query("limit")
	fromDate := ctx.Query("from")
	toDate := ctx.Query("to")
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


	payments, err := app.getExpenses(client, fromDate, toDate, limitNum)
	if err != nil {
		err = fmt.Errorf("unable to get expenses: %v", err)
	}
	app.Logger.Debugf("error occured: %v", err)
	HandleResultOrError(ctx, gin.H{"test": payments, "errorMsg": err}, err, "")
}

func(app *expensesApp) getExpenses(client expensespb.ExpensesServiceClient, fromDate, toDate string, limit int) ([]*expensespb.Payment, error) {
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

	app.Logger.Debugf("pidr fromdate %v todate %v", unixFromDate, unixToDate)
	req := &expensespb.ExpensesRequest{
		FromDate:unixFromDate,
		ToDate:unixToDate,
		Limit:int32(limit),
	}

	res, err := client.GetExpenses(ctx, req)
	app.Logger.Debugf("pidr res %v err %v", res, err)
	if err != nil {
		return nil, err
	}
	return res.Payments, nil
}