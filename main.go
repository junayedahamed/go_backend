package main

import (
	"fmt"
	"go_api/internal/middleware"
	"log"
	"os"

	"strconv"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Greets struct {
	UserName string `json:"username"`
	UserId   int    `json:"userId"`
	Message  string `json:"message"`
}

type Config struct {
	Port string `env:"PORT"`
}

func main() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
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
	// cors

	// router.Use(
	// 	middleware.CORS(),
	// )

	router.Use(middleware.Logger())

	//
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

		ctx.JSON(200, gin.H{
			"messages": "success",
			"data":     messages,
		})

	})

	// search using all
	router.GET("/search-user", func(ctx *gin.Context) {
		q := ctx.Query("q")
		if q == "" {
			ctx.JSON(400, gin.H{
				"message": "No user found",
				"data":    []Greets{},
			})
			return
		}

		qLower := strings.ToLower(q)
		id, err := strconv.Atoi(q)
		hasID := err == nil

		var findData []Greets

		for _, m := range messages {
			userIdStr := strconv.Itoa(m.UserId)
			if strings.Contains(strings.ToLower(m.UserName), qLower) ||
				strings.Contains(strings.ToLower(m.Message), qLower) ||
				strings.Contains(userIdStr, q) ||
				(hasID && m.UserId == id) {
				findData = append(findData, m)
			}
		}

		if len(findData) > 0 {
			ctx.JSON(200, gin.H{
				"message": "success",
				"data":    findData,
			})
		} else {
			ctx.JSON(400, gin.H{
				"message": "No user found",
				"data":    []Greets{},
			})
		}
	})

	router.POST("/add-person", func(ctx *gin.Context) {

		var greets Greets
		if err := ctx.BindJSON(&greets); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		messages = append(messages, greets)
		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    greets,
		})

	})

	//  make auth

	router.POST("/login", func(ctx *gin.Context) {

		var loginUser Greets

		if err := ctx.BindJSON(&loginUser); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		if loginUser.UserName == "" {
			ctx.JSON(400, gin.H{
				"message": "username is required",
			})
			return
		}
		if loginUser.UserId == 0 {
			ctx.JSON(
				400,
				gin.H{
					"message": "user id is required",
				},
			)
			return
		}

		for _, user := range messages {
			if user.UserId == loginUser.UserId && user.UserName == loginUser.UserName {
				ctx.JSON(200, gin.H{
					"message": "success",
					"data":    user,
				})
				return
			}
		}

		ctx.JSON(401, gin.H{
			"message": "Invalid credantial",
		})

	})

	// put user by id
	router.PUT("/update-user/:id", func(ctx *gin.Context) {

		var updateUser Greets
		id := ctx.Params.ByName("id")
		parsedId, err := strconv.Atoi(id)

		if err != nil {
			ctx.JSON(303, gin.H{
				"message": "enter correct id",
			})
			return
		}
		if err := ctx.ShouldBind(&updateUser); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		if updateUser.UserName == "" && updateUser.Message == "" {
			ctx.JSON(400, gin.H{
				"message": "please provide at least one field to update (username or message)",
			})
			return
		}

		for i, user := range messages {
			if user.UserId == parsedId {
				if updateUser.UserName != "" {
					user.UserName = updateUser.UserName
				}
				if updateUser.Message != "" {
					user.Message = updateUser.Message
				}
				messages[i] = user
				ctx.JSON(200, gin.H{
					"message": "success",
					"data":    messages[i],
				})
				return
			}

		}

	})

	// delete user by id

	router.DELETE("/del-user/:id", func(ctx *gin.Context) {

		id := ctx.Params.ByName("id")

		uid, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(303, gin.H{
				"messages": "enter correct id",
			})
			return
		}

		for i, m := range messages {
			if uid == m.UserId {
				messages = append(messages[:i], messages[i+1:]...)
				ctx.JSON(200, gin.H{
					"messages": "success",
					"data":     messages,
				})
				return
			}
		}

		ctx.JSON(
			404,
			gin.H{
				"messages": "Item Not found",
			},
		)

	})

	var config Config

	// 2. Attempt to load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: Error loading .env file, relying on system vars or defaults:", err)
	}

	// 3. Parse variables from the environment into the config struct
	if err := env.Parse(&config); err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}

	router.Run(":" + config.Port)

}
