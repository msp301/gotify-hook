package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SlackMessage struct {
	Text string `json:"text" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/hooks/slack", func(ctx *gin.Context) {
		var msg SlackMessage
		if err := ctx.ShouldBindJSON(&msg); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid_payload"})
			return
		}

		fmt.Printf("Got: '%s'\n", msg.Text)
	})

	router.Run(":8080")
}
