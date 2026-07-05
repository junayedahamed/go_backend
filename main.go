package main

import (
	"github.com/gin-gonic/gin"
)

type Greets struct {
	UserName string `json:"username"`
	userId   int    `json:"userId"`
	Message  string `json:"message"`
}

func main() {

	router := gin.Default()

	var messages = []Greets{
		{
			UserName: "Ajoy",
			userId:   1,
			Message:  "Hello ",
		},
		{
			UserName: "Junayed",
			userId:   2,
			Message:  "Hii ",
		},
		{
			UserName: "Murad",
			userId:   3,
			Message:  "Hey ",
		},
	}
	router.GET("/get-data", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"messages": "success",
			"data":     messages,
		})

	})

	router.Run(":3333")

}
