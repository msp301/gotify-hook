package main

import (
	"fmt"
	"gotify-hook/gotify"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SlackMessage struct {
	Text string `json:"text" binding:"required"`
}

func main() {
	server := os.Getenv("GOTIFY_SERVER")
	token := os.Getenv("APP_TOKEN")

	fmt.Printf("'%s'\n", token)

	gotify := gotify.New(server, token)
	router := gin.Default()

	router.POST("/hooks/slack", func(ctx *gin.Context) {
		var msg SlackMessage
		if err := ctx.ShouldBindJSON(&msg); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid_payload"})
			return
		}

		fmt.Printf("Got: '%s'\n", msg.Text)

		gotify.Message("Test", msg.Text, 5)
	})

	router.Run(":8080")
}
