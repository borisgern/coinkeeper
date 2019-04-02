package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/coinkeeper/services"
)

func(app *expensesApp) newRouter(config *services.HTTPServerConfig, log *services.CtxLogger) *gin.Engine {
	if config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ctxLog := log.NewPrefix("router")
	router := gin.New()

	router.RedirectTrailingSlash = false
	router.Use(services.RequestIDMiddleware(), services.LogMiddleware(ctxLog))
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
	app.Logger.Debugf("test")
	ctx.JSON(http.StatusOK, gin.H{"test": "errorCode"})
}