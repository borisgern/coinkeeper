package services

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")

		if requestID == "" {
			uuid4, _ := uuid.NewV4()
			requestID = uuid4.String()
		}
		c.Set("RequestId", requestID)

		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}

func LogMiddleware(log *CtxLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.WithFields(logrus.Fields{
			"request-id": c.GetString("RequestId"),
		}).Infof("%v: %v from: %v", c.Request.Method, c.Request.URL, c.Request.RemoteAddr)
		c.Next()
	}
}

func HandleResultOrError(ctx *gin.Context, data interface{}, err error, errorCode string) {
	if err == nil {
		Success(ctx, data)
	} else {
		ErrorBadRequest(ctx, err, errorCode)
	}
}

func Success(ctx *gin.Context, obj interface{}) {
	ctx.JSON(200, obj)
}

// ErrorBadRequest - send 400 response with error message
func ErrorBadRequest(ctx *gin.Context, err error, errorCode string) {
	httpError(http.StatusBadRequest, ctx, err, errorCode)
}

func httpError(status int, ctx *gin.Context, err error, errorCode string) {
	ctx.Abort()
	ctx.JSON(status, gin.H{"error": err.Error(), "errorCode": errorCode})
	// log.Errorf("URL: %s, Status: %v, Error: %s, Error code: %v", ctx.Request.URL, status, err.Error(), errorCode)
}

func ErrorForbidden(ctx *gin.Context, err error, errorCode string) {
	httpError(http.StatusForbidden, ctx, err, errorCode)
}