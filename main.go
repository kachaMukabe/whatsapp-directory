package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func verifyWebhook(c *gin.Context) {
	mode := c.Query("hub.mode")
	challenge := c.Query("hub.challenge")
	verifyToken := c.Query("hub.verify_token")
	token := os.Getenv("VERIFYTOKEN")

	if mode == "subscribe" && verifyToken == token {
		c.String(http.StatusOK, "%s", challenge)
	}
}

func handleMessageNotification(c *gin.Context) {
	var notification Notification
	if err := c.BindJSON(&notification); err != nil {
		return
	}
	fmt.Println(notification)
	fmt.Printf("%#v\n", notification)

	c.JSON(http.StatusOK, notification.Entry[0].Changes[0].Value.Messages[0].Text.Body)

}

func main() {
	router := gin.Default()
	router.GET("/webhook", verifyWebhook)
	router.POST("/webhook", handleMessageNotification)

	router.Run("localhost:8000")
}
