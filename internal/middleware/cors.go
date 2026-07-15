package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins: []string{
				"http://localhost:8080",
			},
			MaxAge: 12 * time.Hour,
			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"PATCH",
			},

			AllowHeaders: []string{
				"Origin",
				"Content-Type",
				"Authorization",
			},
		},
	)
}
