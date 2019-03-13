package main

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//if GinReleaseMode {
	//	gin.SetMode(gin.ReleaseMode)
	//}
	router := gin.New()

	router.RedirectTrailingSlash = false
	//router.Use(RequestIDMiddleware(), LogMiddleware(ctxLog))

	groupPayment := router.Group("/statistic")
	groupPayment.GET("/", )

	return router
}