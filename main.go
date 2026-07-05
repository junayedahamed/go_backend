package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Greets struct {
	UserName string `json:"username"`
	UserId   int    `json:"userId"`
	Message  string `json:"message"`
}

func main() {

	router := gin.Default()

	var messages = []Greets{
		{
			UserName: "Ajoy",
			UserId:   1,
			Message:  "Hello ",
		},
		{
			UserName: "Junayed",
			UserId:   2,
			Message:  "Hii ",
		},
		{
			UserName: "Murad",
			UserId:   3,
			Message:  "Hey ",
		},
	}
	router.GET("/get-data", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"messages": "success",
			"data":     messages,
		})

	})

	router.GET("/get-user-by-id/:id", func(ctx *gin.Context) {

		id := ctx.Params.ByName("id")
		parsedId, err := strconv.Atoi(id)

		if err != nil {

			ctx.JSON(303, gin.H{
				"messages": "Enter correct id",
			})
			return
		}

		for _, msg := range messages {

			if msg.UserId == parsedId {

				ctx.JSON(200, gin.H{
					"message": "success",
					"data":    msg,
				})
				return
			}

		}

		ctx.JSON(400, gin.H{
			"messages": "No user found",
		})

	})

	// search using all
	router.GET("/search-user", func(ctx *gin.Context) {

		userName := ctx.Query("username")
		id, err := strconv.Atoi(ctx.Query("id"))
		message := ctx.Query("message")

		var findData []Greets

		if err != nil {

			ctx.JSON(303, gin.H{
				"message": "Enter correct id",
			})
			return

		}

		if userName != "" || id > math.MinInt && id < math.MaxInt || message != "" {

			for _, m := range messages {

				if strings.EqualFold(m.UserName, userName) || m.UserId == id || strings.EqualFold(message, m.Message) {
					findData = append(findData, m)
				}
			}

			ctx.JSON(200, gin.H{
				"message": "success",
				"data":    findData,
			})
			return

		}
		ctx.JSON(400, gin.H{
			"message": "No user found",
		})

	})

	router.Run(":3333")

}
