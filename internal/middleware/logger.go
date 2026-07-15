package middleware

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		println("Request:", ctx.Request.Method, ctx.Request.URL.Path)

		ctx.Next()

		println("Status:", ctx.Writer.Status())
	}
}
