package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Greets struct {
	Message string `json:"message"`
}

func main() {

	router := gin.Default()

	router.GET("/ping/:num1/:num2", func(c *gin.Context) {
		fmt.Print("here")
		num1, err1 := strconv.ParseFloat(c.Param("num1"), 64)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid params",
			})
			return

		}

		num2, err2 := strconv.ParseFloat(c.Param("num2"), 64)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid params",
			})
			return

		}

		sum := num1 + num2
		c.JSON(200, gin.H{
			"message": "pong",
			"sum":     sum,
		})
	})

	router.GET("/sub/:num1/:num2", func(c *gin.Context) {
		num1, err1 := strconv.ParseFloat(c.Param("num1"), 64)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid params",
			})
			return

		}

		num2, err2 := strconv.ParseFloat(c.Param("num2"), 64)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid params",
			})
			return

		}

		sum := num1 - num2
		c.JSON(200, gin.H{

			"sum": sum,
		})

	})
	// fmt.Println("http://localhost:3030")

	router.GET("/item", func(c *gin.Context) {
		query := c.Query("q")
		fmt.Println(query)
		data := []Greets{
			{Message: "hi"},
			{Message: "Hello"},
			{Message: "Test"},
		}

		var results []Greets
		query = strings.ToLower(query)
		fmt.Print(query)
		for _, d := range data {
			message := strings.ToLower(d.Message)
			fmt.Println(strings.Contains(message, query))
			if strings.HasPrefix(message, query) || strings.HasSuffix(message, query) || strings.Contains(message, query) {
				// result = append(result, product)

				results = append(results, d)
			}
		}
		c.JSON(200, gin.H{
			"data": results,
		})

	})

	type user struct {
		uid      string `json:"uid"`
		password string `json:"password"`
	}

	var users = []user{
		{uid: "admin", password: "1234"},
		{uid: "user", password: "[PASSWORD]"},
	}

	router.GET("/login-user/:uid/:pass", func(ctx *gin.Context) {
		uid := ctx.Params.ByName("uid")
		pass := ctx.Params.ByName("pass")

		for i := 0; i < len(users); i++ {
			if uid == users[i].uid && pass == users[i].password {
				ctx.JSON(200, gin.H{
					"message": "Login success",
				})
				return
			}
		}

		ctx.JSON(200, gin.H{
			"message": "Login failed",
		})
	})

	router.Run(":3333")

}
